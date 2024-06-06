package entity

import (
	"SkyTicket/pkg/request"
	"context"
	"database/sql"
	"time"
)

type Flight struct {
	Id                   int64     `json:"id,omitempty"`
	Name                 string    `json:"name,omitempty"`
	From                 string    `json:"from,omitempty"`
	To                   string    `json:"to,omitempty"`
	DepartureDate        time.Time `json:"departure_date,omitempty"`
	ArrivalDate          time.Time `json:"arrival_date,omitempty"`
	AvailableFirstSlot   int       `json:"available_first_slot,omitempty"`
	AvailableEconomySlot int       `json:"available_economy_slot,omitempty"`
	Status               string    `json:"status,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
	UpdatedAt            time.Time `json:"updated_at,omitempty"`
}
type FlightRepository struct {
	Db *sql.DB
}

func (m FlightRepository) CreateFlight(ctx context.Context, f *Flight) (*Flight, error) {
	stmt := `INSERT INTO flights (name, from_location, to_location, departure_date, arrival_date, available_first_slot, available_economy_slot, status) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var id int64
	err := m.Db.QueryRowContext(ctx, stmt, f.Id, f.Name, f.From, f.To, f.DepartureDate, f.ArrivalDate, f.AvailableFirstSlot, f.AvailableEconomySlot, f.Status).Scan(&id)
	if err != nil {
		return nil, err
	}

	f.Id = id
	return f, nil
}

func (m FlightRepository) UpdateFlight(ctx context.Context, id int64, f *Flight) (*Flight, error) {
	stmt := `UPDATE flights SET name=$1, from_location=$2, to_location=$3, departure_date=$4, arrival_date=$5, 
				available_first_slot=$6, available_economy_slot=$7, status=$8 WHERE id=$9`

	_, err := m.Db.ExecContext(ctx, stmt, f.Name, f.From, f.To, f.DepartureDate, f.ArrivalDate, f.AvailableFirstSlot, f.AvailableEconomySlot, f.Status, id)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (m FlightRepository) GetFlight(ctx context.Context, id int64) (*Flight, error) {
	stmt := `SELECT id, name, from_location, to_location, departure_date, arrival_date, 
				available_first_slot, available_economy_slot, status FROM flights WHERE id=$1`

	var f Flight
	err := m.Db.QueryRowContext(ctx, stmt, id).Scan(&f.Id, &f.Name, &f.From, &f.To, &f.DepartureDate, &f.ArrivalDate, &f.AvailableFirstSlot, &f.AvailableEconomySlot, &f.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &f, nil
}

func (m FlightRepository) GetFlightById(ctx context.Context, id int64) (*Flight, error) {
	stmt := `SELECT id, name, from_location, to_location, departure_date, arrival_date, available_first_slot, available_economy_slot, status FROM flights WHERE id=$1`

	var f Flight
	err := m.Db.QueryRowContext(ctx, stmt, id).Scan(&f.Id, &f.Name, &f.From, &f.To, &f.DepartureDate, &f.ArrivalDate, &f.AvailableFirstSlot, &f.AvailableEconomySlot, &f.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	f.Id = id
	return &f, nil
}

func (m FlightRepository) ListFlight(ctx context.Context, paging request.Paging) ([]*Flight, request.Paging, error) {
	stmt := `SELECT id, name, from_location, to_location, departure_date, arrival_date, 
				available_first_slot, available_economy_slot, status FROM flights OFFSET $1 LIMIT $2`

	rows, err := m.Db.QueryContext(ctx, stmt, (paging.Page-1)*paging.Limit, paging.Limit)
	if err != nil {
		return nil, paging, err
	}
	defer rows.Close()

	var flights []*Flight
	for rows.Next() {
		var f Flight
		err := rows.Scan(&f.Id, &f.From, &f.To, &f.DepartureDate, &f.ArrivalDate, &f.AvailableFirstSlot, &f.AvailableEconomySlot, &f.Status)
		if err != nil {
			return nil, paging, err
		}
		flights = append(flights, &f)
	}

	if err := rows.Err(); err != nil {
		return nil, paging, err
	}

	totalRowsStmt := `SELECT count(*) FROM flights`
	var total int64
	err = m.Db.QueryRowContext(ctx, totalRowsStmt).Scan(&total)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = total
	paging.Process()

	return flights, paging, nil
}

func (m FlightRepository) DeleteFlight(ctx context.Context, id int64) error {
	stmt := `DELETE FROM flights WHERE id=$1`
	_, err := m.Db.ExecContext(ctx, stmt, id)
	return err
}

func (m FlightRepository) SearchFlight(ctx context.Context, from, to string, departureDate, arrivalDate time.Time, paging request.Paging) ([]*Flight, request.Paging, error) {
	stmt := `SELECT id, name, from_location, to_location, departure_date, arrival_date, 
				available_first_slot, available_economy_slot, status FROM flights 
				WHERE from_location=$1 AND to_location=$2 AND departure_date >= $3 AND arrival_date <= $4 AND status=$5
				OFFSET $6 LIMIT $7`

	rows, err := m.Db.QueryContext(ctx, stmt, from, to, departureDate, arrivalDate, (paging.Page-1)*paging.Limit, paging.Limit)
	if err != nil {
		return nil, paging, err
	}
	defer rows.Close()

	var flights []*Flight
	for rows.Next() {
		var f Flight
		err := rows.Scan(&f.Id, &f.From, &f.To, &f.DepartureDate, &f.ArrivalDate, &f.AvailableFirstSlot, &f.AvailableEconomySlot, &f.Status)
		if err != nil {
			return nil, paging, err
		}
		flights = append(flights, &f)
	}

	if err := rows.Err(); err != nil {
		return nil, paging, err
	}

	totalRowsStmt := `SELECT count(*) FROM flights WHERE from_location=$1 AND to_location=$2 AND departure_date >= $3 AND arrival_date <= $4 AND status=$5`
	var total int64
	err = m.Db.QueryRowContext(ctx, totalRowsStmt, from, to, departureDate, arrivalDate).Scan(&total)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = total
	paging.Process()

	return flights, paging, nil
}
