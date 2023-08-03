package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	AppPort string
}

type Config struct {
	ApiConfig
}

func (c Config) readConfigFile() Config {
	c.ApiConfig = ApiConfig{
		AppPort: os.Getenv("APP_PORT"),
	}
	return c
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := Config{}
	return cfg.readConfigFile()
}
