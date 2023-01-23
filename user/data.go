package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"autoIncrement:11"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Gender    string         `json:"gender"`
	Address   string         `json:"address"`
	Dob       string         `json:"dob"`
	Image     string         `json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
