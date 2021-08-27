package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          uuid.UUID `json:"id" gorm:"type=uuid;default:uuid_generate_v4()"`
	Slut        string    `json:"slut" gorm:"type:varchar(20);not null;unique"`
	Name        string    `json:"name" gorm:"type:varchar(256);not null"`
	Address     string    `json:"address" gorm:"type:varchar(256);"`
	LicenseId   string    `json:"license_id" gorm:"type:varchar(256);"`
	PhoneNumber string    `json:"phone" gorm:"type:bigint;unique;not null;"`
	Email       string    `json:"email" gorm:"type:varchar(256);not null;unique;"`
	Password    string    `json:"password" gorm:"type:varchar(256)"`
	Active      bool      `json:"active" gorm:"type:bool"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type History struct {
}
