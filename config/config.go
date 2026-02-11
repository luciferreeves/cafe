package config

import (
	"cafe/utils/env"
	"log"

	"github.com/joho/godotenv"
)

var (
	Server   server
	Database database
	Session  session
	OpenID   openid
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

	if err := env.Parse(&Session); err != nil {
		log.Fatalf("Failed to parse SessionConfig: %v", err)
	}

	if err := env.Parse(&OpenID); err != nil {
		log.Fatalf("Failed to parse OpenIDConfig: %v", err)
	}
}
