package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envFile := "config/.env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
