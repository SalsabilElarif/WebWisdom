package config

import (
	"log"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}