package ticket

import (
	"SkyTicket/internal/domain/entity"
	"context"
)

type UseCase interface {
	BookTicket(ctx context.Context) (*entity.Ticket, error)
}

type TicketUseCase struct {
	data entity.TicketData
}

func NewTicketUseCase(data entity.TicketData) *TicketUseCase {
	return &TicketUseCase{data: data}
}

func (uc *TicketUseCase) BookTicket(ctx context.Context) (*entity.Ticket, error) {
	return nil, nil
}
