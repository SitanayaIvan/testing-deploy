package database

import (
	"github.com/SitanayaIvan/testing-deploy/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	conf, _ := config.LoadVariable()

	db, err := gorm.Open(postgres.Open(conf.Db.DBUrl), &gorm.Config{SkipDefaultTransaction: true})

	return db, err
}
