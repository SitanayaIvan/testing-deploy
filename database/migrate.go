package database

import (
	"github.com/SitanayaIvan/testing-deploy/user"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
