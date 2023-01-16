package database

import (
	"os/user"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
