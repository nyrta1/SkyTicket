package data

import (
	"database/sql"
)

// DBConnection handles the database connection
type DBConnection struct {
	DB *sql.DB
}

func NewDBConnection(connectionString string) (*DBConnection, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &DBConnection{DB: db}, nil
}
