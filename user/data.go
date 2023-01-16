package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Dob       string `json:"dob"`
	Image     string `json:"image"`
}
