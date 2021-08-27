package requests

import (
	"github.com/google/uuid"
)

type BookingRequest struct {
	Id         uuid.UUID `json:"id"`
	Slut       string    `json:"slut"`
	CustomerId uuid.UUID `json:"customer"`
	FlightId   uuid.UUID `json:"flight"`
}
