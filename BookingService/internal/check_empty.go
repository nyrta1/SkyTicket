package internal

import (
	"SkyTicket/entity"
)

func CheckFlightEmptyInput(flightInput, flight *entity.Flight) *entity.Flight {
	if flightInput.Name == "" {
		flightInput.Name = flight.Name
	}
	if flightInput.From == "" {
		flightInput.From = flight.From
	}
	if flightInput.To == "" {
		flightInput.To = flight.To
	}
	if flightInput.AvailableSlot == 0 {
		flightInput.AvailableSlot = flight.AvailableSlot
	}
	if flightInput.Status == "" {
		flightInput.Status = flight.Status
	}
	return flightInput
}
