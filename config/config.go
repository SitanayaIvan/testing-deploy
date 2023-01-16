package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadVariable() (ConfigBody, error) {
	var conf ConfigBody

	err := godotenv.Load()
	if err != nil {
		log.Println("[Error] loading .env file")
	}

	// DB
	conf.Db.Host = os.Getenv("DB_HOST")
	conf.Db.Name = os.Getenv("DB_NAME")
	conf.Db.User = os.Getenv("DB_USER")
	conf.Db.Password = os.Getenv("DB_PASSWORD")
	conf.Db.Port = os.Getenv("DB_PORT")
	conf.Db.DBUrl = os.Getenv("DB_URL")

	// APP
	conf.App.Port = os.Getenv("PORT")

	return conf, nil
}
