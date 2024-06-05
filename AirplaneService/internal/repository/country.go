package repository

import (
	"AirplaneService/internal/entity"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (repo *CountryRepository) Add(ctx context.Context, country *entity.Country) error {
	query := `INSERT INTO country (name, prefix) VALUES ($1, $2)`
	_, err := repo.db.ExecContext(ctx, query, country.Name, country.Prefix)
	return err
}

func (repo *CountryRepository) GetById(ctx context.Context, id int64) (*entity.Country, error) {
	query := `SELECT id, name, prefix FROM country WHERE id=$1`
	row := repo.db.QueryRowContext(ctx, query, id)

	var country entity.Country
	err := row.Scan(&country.Id, &country.Name, &country.Prefix)
	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (repo *CountryRepository) GetAll(ctx context.Context) ([]*entity.Country, error) {
	query := `SELECT id, name, prefix FROM country`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []*entity.Country
	for rows.Next() {
		var country entity.Country
		err := rows.Scan(&country.Id, &country.Name, &country.Prefix)
		if err != nil {
			return nil, err
		}
		countries = append(countries, &country)
	}

	return countries, nil
}

func (repo *CountryRepository) UpdateById(ctx context.Context, id int64, country *entity.Country) error {
	query := `UPDATE country SET name=$1, prefix=$2 WHERE id=$3`
	_, err := repo.db.ExecContext(ctx, query, country.Name, country.Prefix, id)
	return err
}

func (repo *CountryRepository) DeleteById(ctx context.Context, id int64) error {
	query := `DELETE FROM country WHERE id=$1`
	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}
