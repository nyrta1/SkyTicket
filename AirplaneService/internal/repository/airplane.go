package repository

import (
	"AirplaneService/internal/entity"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type AirplaneRepository struct {
	db *sql.DB
}

func NewAirplaneRepository(db *sql.DB) *AirplaneRepository {
	return &AirplaneRepository{db: db}
}

func (repo *AirplaneRepository) Add(ctx context.Context, airplane *entity.Airplane) error {
	query := `INSERT INTO airplane (manufacturer_id, manufacturer_year, first_slot_capacity, 
		economy_slot_capacity, country_origin_id, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.ExecContext(ctx, query, airplane.ManufacturerId, airplane.ManufacturerYear,
		airplane.FirstSlotCapacity, airplane.EconomySlotCapacity, airplane.CountryOriginId, time.Now())
	return err
}

func (repo *AirplaneRepository) GetById(ctx context.Context, id int64) (*entity.Airplane, error) {
	query := `SELECT id, manufacturer_id, manufacturer_year, first_slot_capacity, economy_slot_capacity, country_origin_id FROM airplane WHERE id = $1`
	row := repo.db.QueryRowContext(ctx, query, id)

	var airplane entity.Airplane
	err := row.Scan(&airplane.Id, &airplane.ManufacturerId, &airplane.ManufacturerYear,
		&airplane.FirstSlotCapacity, &airplane.EconomySlotCapacity, &airplane.CountryOriginId)
	if err != nil {
		return nil, err
	}

	return &airplane, nil
}

func (repo *AirplaneRepository) GetAll(ctx context.Context) ([]*entity.Airplane, error) {
	query := `SELECT id, manufacturer_id, manufacturer_year, first_slot_capacity, economy_slot_capacity, country_origin_id FROM airplane`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airplanes []*entity.Airplane
	for rows.Next() {
		var airplane entity.Airplane
		err := rows.Scan(&airplane.Id, &airplane.ManufacturerId, &airplane.ManufacturerYear,
			&airplane.FirstSlotCapacity, &airplane.EconomySlotCapacity, &airplane.CountryOriginId)
		if err != nil {
			return nil, err
		}
		airplanes = append(airplanes, &airplane)
	}

	return airplanes, nil
}

func (repo *AirplaneRepository) UpdateById(ctx context.Context, id int64, airplane *entity.Airplane) error {
	query := `UPDATE airplane SET manufacturer_id=$1, manufacturer_year=$2, first_slot_capacity=$3, 
		economy_slot_capacity=$4, country_origin_id=$5, updated_at=$6 WHERE id=$7`
	_, err := repo.db.ExecContext(ctx, query, airplane.ManufacturerId, airplane.ManufacturerYear,
		airplane.FirstSlotCapacity, airplane.EconomySlotCapacity, airplane.CountryOriginId, time.Now(), id)
	return err
}

func (repo *AirplaneRepository) DeleteById(ctx context.Context, id int64) error {
	query := `DELETE FROM airplane WHERE id=$1`
	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}
