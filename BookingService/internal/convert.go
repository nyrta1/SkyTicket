package internal

import (
	"strings"
)

// Status enumeration for flight status
type Status string

const (
	StatusAvailable Status = "available"
	StatusArrived   Status = "arrived"
	StatusCancel    Status = "cancel"
)

var flightStatus = map[string]Status{
	"available": StatusAvailable,
	"arrived":   StatusArrived,
	"cancel":    StatusCancel,
}

// ParseString parses a string and returns the corresponding Status enum value
func ParseString(str string) (Status, bool) {
	c, ok := flightStatus[strings.ToLower(str)]
	return c, ok
}
