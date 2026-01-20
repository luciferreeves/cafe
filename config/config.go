package config

import (
	"cafe/utils/env"
	"log"

	"github.com/joho/godotenv"
)

var (
	Server   server
	Database database
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	if err := env.Parse(&Server); err != nil {
		log.Fatalf("Failed to parse ServerConfig: %v", err)
	}

	if err := env.Parse(&Database); err != nil {
		log.Fatalf("Failed to parse DatabaseConfig: %v", err)
	}
}
