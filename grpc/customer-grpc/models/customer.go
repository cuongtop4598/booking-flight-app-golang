package models

type Customer struct {
	Id          string `json:"id" gorm:"type=uuid;default:uuid_generate_v4()"`
	Name        string `json:"name" gorm:"type:varchar(256);not null"`
	Address     string `json:"address" gorm:"type:varchar(256);"`
	LicenseId   string `json:"license_id" gorm:"type:varchar(256);"`
	PhoneNumber string `json:"phone" gorm:"type:interger"`
	Email       string `json:"email" gorm:"type:varchar(256);not null;unique;"`
	Password    string `json:"password" gorm:"type:varchar(256)"`
	Active      bool   `json:"active" gorm:"type:varchar(256)"`
}
