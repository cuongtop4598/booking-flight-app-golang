package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	Id          uuid.UUID `json:"id" gorm:"type=uuid;default:uuid_generate_v4();"`
	Slut        string    `json:"slut" gorm:"type:varchar(20);not null;unique"`
	CustomerId  uuid.UUID `json:"customer_id" gorm:"type=uuid"`
	FlightId    uuid.UUID `json:"flight_id" gorm:"type=uuid"`
	Code        string    `json:"code" gorm:"type:varchar(256); unique"`
	Status      string    `json:"status" gorm:"type:varchar(20);"`
	Booked_date time.Time `json:"booked_date"`
	CreatedAt   time.Time `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
