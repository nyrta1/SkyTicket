package repository

import (
	"SkyTicket/entity"
	"context"
	"database/sql"
	"time"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, b *entity.Booking) (*entity.Booking, error)
	GetTicketClass(ctx context.Context, id int64) (*entity.Ticket, error)
	UpdateBooking(ctx context.Context, id int64, b *entity.Booking) (*entity.Booking, error)
	GetBookingByCode(ctx context.Context, code string) (*entity.Booking, error)
	GetBookingHistory(ctx context.Context, customerId int64) ([]*entity.Booking, error)
	GetBookingByFlight(ctx context.Context, flightId int64) ([]*entity.Booking, error)
	UpdateBookingStatus(ctx context.Context, id int64, status string) error
	ListBooking(ctx context.Context) ([]*entity.Booking, error)
}

type FlightRepository interface {
	CreateFlight(ctx context.Context, f *entity.Flight) (*entity.Flight, error)
	GetFlight(ctx context.Context, name string) (*entity.Flight, error)
	ListFlight(ctx context.Context) ([]*entity.Flight, error)
	UpdateFlight(ctx context.Context, id int64, f *entity.Flight) (*entity.Flight, error)
	DeleteFlight(ctx context.Context, id int64) error
	GetFlightById(ctx context.Context, id int64) (*entity.Flight, error)
	SearchFlight(ctx context.Context, from, to string, departureDate, arrivalDate time.Time) ([]*entity.Flight, error)
}

type Models struct {
	Booking entity.BookingModel
	Flight  entity.FlightModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Booking: entity.BookingModel{Db: db},
		Flight:  entity.FlightModel{Db: db},
	}
}
