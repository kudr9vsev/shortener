package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
}

func LoadDBConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error load .env file, use default conf")
	}

	return &Config{
		DSN: os.Getenv("DSN"),
	}
}
