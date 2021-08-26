package requests

import (
	"time"

	"github.com/gofrs/uuid"
)

type UpdateFlightRequest struct {
	Id            uuid.UUID
	Slut          string
	Name          string
	From          string
	To            string
	Date          time.Time
	Status        string
	AvailableSlot int32
}

type ListFlightRequest struct {
	Name string
	From string
	To   string
	Date time.Time
}
