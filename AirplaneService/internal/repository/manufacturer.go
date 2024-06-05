package repository

import (
	"AirplaneService/internal/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ManufacturerRepository struct {
	Db *sql.DB
}

func NewManufacturerRepository(db *sql.DB) *ManufacturerRepository {
	return &ManufacturerRepository{Db: db}
}

func (r *ManufacturerRepository) Add(ctx context.Context, entity *entity.Manufacturer) error {
	sqlQuery := "INSERT INTO manufacturer (name) VALUES ($1)"
	_, err := r.Db.ExecContext(ctx, sqlQuery, entity.Name)
	if err != nil {
		log.Println("Error adding manufacturer:", err)
		return err
	}
	return nil
}

func (r *ManufacturerRepository) GetById(ctx context.Context, id int64) (*entity.Manufacturer, error) {
	sqlQuery := "SELECT id, name FROM manufacturer WHERE id=$1"
	row := r.Db.QueryRowContext(ctx, sqlQuery, id)
	var manufacturer entity.Manufacturer
	err := row.Scan(&manufacturer.Id, &manufacturer.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Println("Error getting manufacturer by ID:", err)
		return nil, err
	}
	return &manufacturer, nil
}

func (r *ManufacturerRepository) GetAll(ctx context.Context) ([]*entity.Manufacturer, error) {
	sqlQuery := "SELECT id, name FROM manufacturer"
	rows, err := r.Db.QueryContext(ctx, sqlQuery)
	if err != nil {
		log.Println("Error getting all manufacturers:", err)
		return nil, err
	}
	defer rows.Close()
	fmt.Println("1", rows)

	var manufacturers []*entity.Manufacturer
	for rows.Next() {
		var manufacturer entity.Manufacturer
		err := rows.Scan(&manufacturer.Id, &manufacturer.Name)
		if err != nil {
			log.Println("Error scanning manufacturer:", err)
			return nil, err
		}
		manufacturers = append(manufacturers, &manufacturer)
	}
	fmt.Println("2", manufacturers)

	if err = rows.Err(); err != nil {
		log.Println("Error with rows during getting all manufacturers:", err)
		return nil, err
	}

	return manufacturers, nil
}

func (r *ManufacturerRepository) UpdateById(ctx context.Context, id int64, entity *entity.Manufacturer) error {
	sqlQuery := "UPDATE manufacturer SET name=$1 WHERE id=$2"
	_, err := r.Db.ExecContext(ctx, sqlQuery, entity.Name, id)
	if err != nil {
		log.Println("Error updating manufacturer by ID:", err)
		return err
	}
	return nil
}

func (r *ManufacturerRepository) DeleteById(ctx context.Context, id int64) error {
	sqlQuery := "DELETE FROM manufacturer WHERE id=$1"
	_, err := r.Db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		log.Println("Error deleting manufacturer by ID:", err)
		return err
	}
	return nil
}
