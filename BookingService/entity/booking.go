package entity

import (
	"context"
	"database/sql"
	"time"
)

type Booking struct {
	ID         int64     `json:"id,omitempty"`
	CustomerID int64     `json:"customer_id,omitempty"`
	FlightID   int64     `json:"flight_id,omitempty"`
	Code       string    `json:"code,omitempty"`
	Status     string    `json:"status,omitempty"`
	TicketID   int64     `json:"ticket_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type BookingModel struct {
	Db *sql.DB
}
type Ticket struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (r *BookingModel) GetTicketClass(ctx context.Context, id int64) (*Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BookingModel) CreateBooking(ctx context.Context, b *Booking) (*Booking, error) {
	query := `INSERT INTO bookings (code, customer_id, flight_id, status, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.Db.QueryRowContext(ctx, query, b.Code, b.CustomerID, b.FlightID, b.Status, b.CreatedAt, b.UpdatedAt).Scan(&b.ID)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *BookingModel) UpdateBooking(ctx context.Context, id int64, b *Booking) (*Booking, error) {
	query := `UPDATE bookings SET code = $1, customer_id = $2, flight_id = $3, status = $4, updated_at = $5 WHERE id = $6`
	_, err := r.Db.ExecContext(ctx, query, b.Code, b.CustomerID, b.FlightID, b.Status, b.UpdatedAt, id)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *BookingModel) GetBookingByCode(ctx context.Context, code string) (*Booking, error) {
	var booking Booking
	query := `SELECT id, code, customer_id, flight_id, status, created_at, updated_at FROM bookings WHERE code = $1`
	err := r.Db.QueryRowContext(ctx, query, code).Scan(&booking.ID, &booking.Code, &booking.CustomerID, &booking.FlightID, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *BookingModel) GetBookingHistory(ctx context.Context, customerId int64) ([]*Booking, error) {
	query := `SELECT id, code, customer_id, flight_id, status, created_at, updated_at FROM bookings WHERE customer_id = $1`
	rows, err := r.Db.QueryContext(ctx, query, customerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.ID, &booking.Code, &booking.CustomerID, &booking.FlightID, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}
	return bookings, nil
}

func (r *BookingModel) GetBookingByFlight(ctx context.Context, flightId int64) ([]*Booking, error) {
	query := `SELECT id, code, customer_id, flight_id, status, created_at, updated_at FROM bookings WHERE flight_id = $1`
	rows, err := r.Db.QueryContext(ctx, query, flightId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.ID, &booking.Code, &booking.CustomerID, &booking.FlightID, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}
	return bookings, nil
}

func (r *BookingModel) UpdateBookingStatus(ctx context.Context, id int64, status string) error {
	query := `UPDATE bookings SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.Db.ExecContext(ctx, query, status, time.Now(), id)
	return err
}

func (r *BookingModel) ListBooking(ctx context.Context) ([]*Booking, error) {
	query := `SELECT id, code, customer_id, flight_id, status, created_at, updated_at FROM bookings`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.ID, &booking.Code, &booking.CustomerID, &booking.FlightID, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}
	return bookings, nil
}
