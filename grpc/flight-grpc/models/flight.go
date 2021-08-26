package models

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id            uuid.UUID `gorm:"type=uuid;default:uuid_generate_v4()"`
	Slut          string    `gorm:"type:varchar(256);not null;unique"`
	Name          string    `gorm:"type:varchar(256);not null;"`
	From          string    `gorm:"type:varchar(256);not null;"`
	To            string    `gorm:"type:varchar(256);not null;"`
	Date          time.Time
	Status        string `gorm:"type:varchar(50);not null;"`
	AvailableSlot int32  `gorm:"type:integer;not null;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
