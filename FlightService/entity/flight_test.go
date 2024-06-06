package entity

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFlightRepository_CreateFlight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock: %s", err)
	}
	defer db.Close()

	flightRepo := FlightRepository{Db: db}

	flight := &Flight{
		Name:                 "Test Flight",
		From:                 "From Location",
		To:                   "To Location",
		DepartureDate:        time.Now(),
		ArrivalDate:          time.Now().Add(2 * time.Hour),
		AvailableFirstSlot:   50,
		AvailableEconomySlot: 100,
		Status:               "Active",
	}

	mock.ExpectQuery("INSERT INTO flights").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	_, err = flightRepo.CreateFlight(context.Background(), flight)
	assert.NoError(t, err)
}

func TestFlightRepository_UpdateFlight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock: %s", err)
	}
	defer db.Close()

	flightRepo := FlightRepository{Db: db}

	flight := &Flight{
		Name:                 "Test Flight",
		From:                 "From Location",
		To:                   "To Location",
		DepartureDate:        time.Now(),
		ArrivalDate:          time.Now().Add(2 * time.Hour),
		AvailableFirstSlot:   50,
		AvailableEconomySlot: 100,
		Status:               "Active",
	}

	mock.ExpectExec("UPDATE flights").WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = flightRepo.UpdateFlight(context.Background(), 1, flight)
	assert.NoError(t, err)
}

func TestFlightRepository_GetFlight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock: %s", err)
	}
	defer db.Close()

	flightRepo := FlightRepository{Db: db}

	rows := sqlmock.NewRows([]string{"id", "name", "from_location", "to_location", "departure_date", "arrival_date", "available_first_slot", "available_economy_slot", "status"}).
		AddRow(1, "Test Flight", "From Location", "To Location", time.Now(), time.Now().Add(2*time.Hour), 50, 100, "Active")

	mock.ExpectQuery("SELECT id, name, from_location, to_location, departure_date, arrival_date, available_first_slot, available_economy_slot, status FROM flights WHERE id=").
		WithArgs(1).WillReturnRows(rows)

	_, err = flightRepo.GetFlight(context.Background(), 1)
	assert.NoError(t, err)
}

func TestFlightRepository_GetFlightById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock: %s", err)
	}
	defer db.Close()

	flightRepo := FlightRepository{Db: db}

	rows := sqlmock.NewRows([]string{"id", "name", "from_location", "to_location", "departure_date", "arrival_date", "available_first_slot", "available_economy_slot", "status"}).
		AddRow(1, "Test Flight", "From Location", "To Location", time.Now(), time.Now().Add(2*time.Hour), 50, 100, "Active")

	mock.ExpectQuery("SELECT id, name, from_location, to_location, departure_date, arrival_date, available_first_slot, available_economy_slot, status FROM flights WHERE id=").
		WithArgs(1).WillReturnRows(rows)

	_, err = flightRepo.GetFlightById(context.Background(), 1)
	assert.NoError(t, err)
}

func TestFlightRepository_DeleteFlight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock: %s", err)
	}
	defer db.Close()

	flightRepo := FlightRepository{Db: db}

	mock.ExpectExec("DELETE FROM flights").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = flightRepo.DeleteFlight(context.Background(), 1)
	assert.NoError(t, err)
}
