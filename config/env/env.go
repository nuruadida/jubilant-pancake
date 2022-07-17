package env

import (
	"Nurul-trial/models"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
)

var Config models.ServerConfig

func init() {
	err := LoadConfig()
	if err != nil {
		log.Println("cannot open file .env")
	}
}

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		logrus.Warning("config/env : load file env")
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	return err
}
