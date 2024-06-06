package repo

import (
	"SkyTicket/FlightService/entity"
	"SkyTicket/pkg/request"
	"database/sql"
	"golang.org/x/net/context"
	"time"
)

type FlightRepository interface {
	CreateFlight(ctx context.Context, f *entity.Flight) (*entity.Flight, error)
	GetFlight(ctx context.Context, id int64) (*entity.Flight, error)
	ListFlight(ctx context.Context, paging request.Paging) ([]*entity.Flight, request.Paging, error)
	UpdateFlight(ctx context.Context, id int64, f *entity.Flight) (*entity.Flight, error)
	DeleteFlight(ctx context.Context, id int64) error
	GetFlightById(ctx context.Context, id int64) (*entity.Flight, error)
	SearchFlight(ctx context.Context, from, to string, departureDate, arrivalDate time.Time, paging request.Paging) ([]*entity.Flight, request.Paging, error)
}
type Models struct {
	Flight entity.FlightRepository
}

func NewModels(db *sql.DB) Models {
	return Models{
		Flight: entity.FlightRepository{Db: db},
	}
}
