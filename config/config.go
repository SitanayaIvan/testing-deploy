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
	conf.Db.DBUrl = os.Getenv("DB_URL")

	// APP
	conf.App.Port = os.Getenv("PORT")

	return conf, nil
}
