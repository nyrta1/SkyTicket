package entity

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Status string

const (
	StatusAvailable Status = "Available"
	StatusArrived   Status = "Arrived"
	StatusCancel    Status = "Cancel"
)

type Flight struct {
	ID            int64     `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	From          string    `json:"from,omitempty"`
	To            string    `json:"to,omitempty"`
	DepartureDate time.Time `json:"departure_date,omitempty"`
	ArrivalDate   time.Time `json:"arrival_date,omitempty"`
	AvailableSlot int       `json:"available_slot,omitempty"`
	Status        Status    `json:"status,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

func (s Status) String() string {
	return string(s)
}
func StatusValidator(s Status) error {
	switch s {
	case StatusAvailable, StatusArrived, StatusCancel:
		return nil
	default:
		return fmt.Errorf("flight: invalid enum value for status field: %q", s)
	}
}

type FlightModel struct {
	Db *sql.DB
}

func (r *FlightModel) CreateFlight(ctx context.Context, f *Flight) (*Flight, error) {
	query := `INSERT INTO flights (name, from_location, to_location, departure_time, arrival_time, price) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.Db.QueryRowContext(ctx, query, f.Name, f.From, f.To, f.DepartureDate, f.ArrivalDate).Scan(&f.ID)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FlightModel) GetFlight(ctx context.Context, name string) (*Flight, error) {
	var flight Flight
	query := `SELECT id, name, from_location, to_location, departure_time, arrival_time, price FROM flights WHERE name = $1`
	err := r.Db.QueryRowContext(ctx, query, name).Scan(&flight.ID, &flight.Name, &flight.From, &flight.To, &flight.DepartureDate, &flight.ArrivalDate)
	if err != nil {
		return nil, err
	}
	return &flight, nil
}

func (r *FlightModel) ListFlight(ctx context.Context) ([]*Flight, error) {
	query := `SELECT id, name, from_location, to_location, departure_time, arrival_time, price FROM flights`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []*Flight
	for rows.Next() {
		var flight Flight
		err := rows.Scan(&flight.ID, &flight.Name, &flight.From, &flight.To, &flight.DepartureDate, &flight.ArrivalDate)
		if err != nil {
			return nil, err
		}
		flights = append(flights, &flight)
	}
	return flights, nil
}

func (r *FlightModel) UpdateFlight(ctx context.Context, id int64, f *Flight) (*Flight, error) {
	query := `UPDATE flights SET name = $1, from_location = $2, to_location = $3, departure_time = $4, arrival_time = $5, price = $6 WHERE id = $7`
	_, err := r.Db.ExecContext(ctx, query, f.Name, f.From, f.To, f.DepartureDate, f.ArrivalDate, id)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FlightModel) DeleteFlight(ctx context.Context, id int64) error {
	query := `DELETE FROM flights WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

func (r *FlightModel) GetFlightById(ctx context.Context, id int64) (*Flight, error) {
	var flight Flight
	query := `SELECT id, name, from_location, to_location, departure_time, arrival_time, price FROM flights WHERE id = $1`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&flight.ID, &flight.Name, &flight.From, &flight.To, &flight.DepartureDate, &flight.ArrivalDate)
	if err != nil {
		return nil, err
	}
	return &flight, nil
}

func (r *FlightModel) SearchFlight(ctx context.Context, from, to string, departureDate, arrivalDate time.Time) ([]*Flight, error) {
	query := `SELECT id, name, from_location, to_location, departure_time, arrival_time, price 
              FROM flights 
              WHERE from_location = $1 AND to_location = $2 AND departure_time >= $3 AND arrival_time <= $4`
	rows, err := r.Db.QueryContext(ctx, query, from, to, departureDate, arrivalDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []*Flight
	for rows.Next() {
		var flight Flight
		err := rows.Scan(&flight.ID, &flight.Name, &flight.From, &flight.To, &flight.DepartureDate, &flight.ArrivalDate)
		if err != nil {
			return nil, err
		}
		flights = append(flights, &flight)
	}
	return flights, nil
}
