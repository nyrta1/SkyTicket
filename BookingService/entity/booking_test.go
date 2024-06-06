package entity

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBookingModel_GetTicketClass(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingModel := &BookingModel{Db: db}

	ctx := context.Background()
	ticketID := int64(1)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Economy")

	mock.ExpectQuery("SELECT id, name FROM ticket WHERE id = ?").
		WithArgs(ticketID).
		WillReturnRows(rows)

	ticket, err := bookingModel.GetTicketClass(ctx, ticketID)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, int64(1), ticket.ID)
	assert.Equal(t, "Economy", ticket.Name)
}

func TestBookingModel_UpdateBooking(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingModel := &BookingModel{Db: db}

	ctx := context.Background()
	bookingID := int64(1)
	booking := &Booking{
		UserID:    1,
		FlightID:  1,
		Code:      "ABC123",
		Status:    "confirmed",
		TicketID:  1,
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec("UPDATE bookings").
		WithArgs(booking.Code, booking.UserID, booking.FlightID, booking.Status, booking.UpdatedAt, bookingID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	updatedBooking, err := bookingModel.UpdateBooking(ctx, bookingID, booking)
	assert.NoError(t, err)
	assert.NotNil(t, updatedBooking)
}

func TestBookingModel_GetBookingHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingModel := &BookingModel{Db: db}

	ctx := context.Background()
	customerID := int64(1)

	rows := sqlmock.NewRows([]string{"id", "code", "user_id", "flight_id", "status", "created_at", "updated_at"}).
		AddRow(1, "ABC123", 1, 1, "confirmed", time.Now(), time.Now())

	mock.ExpectQuery("SELECT id, code, user_id, flight_id, status, created_at, updated_at FROM booking WHERE user_id = ?").
		WithArgs(customerID).
		WillReturnRows(rows)

	bookings, err := bookingModel.GetBookingHistory(ctx, customerID)
	assert.NoError(t, err)
	assert.NotNil(t, bookings)
	assert.Len(t, bookings, 1)
}

func TestBookingModel_GetBookingByFlight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingModel := &BookingModel{Db: db}

	ctx := context.Background()
	flightID := int64(1)

	rows := sqlmock.NewRows([]string{"id", "code", "user_id", "flight_id", "status", "created_at", "updated_at"}).
		AddRow(1, "ABC123", 1, 1, "confirmed", time.Now(), time.Now())

	mock.ExpectQuery("SELECT id, code, user_id, flight_id, status, created_at, updated_at FROM booking WHERE flight_id = ?").
		WithArgs(flightID).
		WillReturnRows(rows)

	bookings, err := bookingModel.GetBookingByFlight(ctx, flightID)
	assert.NoError(t, err)
	assert.NotNil(t, bookings)
	assert.Len(t, bookings, 1)
}

func TestBookingModel_UpdateBookingStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingModel := &BookingModel{Db: db}

	ctx := context.Background()
	bookingID := int64(1)
	status := "cancelled"

	mock.ExpectExec("UPDATE bookings").
		WithArgs(status, sqlmock.AnyArg(), bookingID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = bookingModel.UpdateBookingStatus(ctx, bookingID, status)
	assert.NoError(t, err)
}
