package entity

import (
	"time"
)

type Ticket struct {
	ID          string
	UserID      string
	FlightID    string
	BookingDate time.Time
}

type MemoryTicketData struct {
	tickets map[string]*Ticket
}

func NewTicketData() *MemoryTicketData {
	return &MemoryTicketData{
		tickets: make(map[string]*Ticket),
	}
}

func (d *MemoryTicketData) Save(ticket Ticket) error {

	return nil
}

func (d *MemoryTicketData) FindByID(id string) (*Ticket, error) {
	return nil, nil
}
