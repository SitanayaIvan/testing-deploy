package database

import (
	"github.com/SitanayaIvan/testing-deploy/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	conf, _ := config.LoadVariable()

	dsn := "host=" + conf.Db.Host + " user=" + conf.Db.User + " password=" + conf.Db.Password + " dbname=" + conf.Db.Name + " port=" + conf.Db.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	return db, err
}
