package config

import (
	"echo-mangosteen/internal/configs"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Config struct {
	DB  configs.Database
	JWT configs.JWT
}

func New() *Config {

	if err := godotenv.Load(); err != nil {
		log.Panicf("falied to load env file: %s", err)
	}

	return &Config{
		DB:  configs.DatabaseStore(),
		JWT: configs.JWTStore(),
	}
}
