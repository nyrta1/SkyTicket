package repository

import (
	"SkyTicket/BookingService/entity"
	"context"
	"database/sql"
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

type Models struct {
	Booking entity.BookingModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Booking: entity.BookingModel{Db: db},
	}
}
