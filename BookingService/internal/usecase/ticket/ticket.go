package ticket

import (
	"SkyTicket/internal/domain/entity"
	"context"
)

type UseCase interface {
	BookTicket(ctx context.Context) (*entity.Ticket, error)
}

type TicketUseCase struct {
	data *entity.MemoryTicketData // Change the type to a pointer
}

func NewTicketUseCase(data *entity.MemoryTicketData) *TicketUseCase { // Change the parameter type
	return &TicketUseCase{data: data}
}

func (uc *TicketUseCase) BookTicket(ctx context.Context) (*entity.Ticket, error) {
	return nil, nil
}
