package main

import (
	"cafe/config"
	"cafe/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: router.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(helmet.New(helmet.Config{
		CrossOriginEmbedderPolicy: "unsafe-none",
	}))
	app.Use(cors.New())

	router.Initialize(app)

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	log.Printf("Starting server at %s\n", address)
	log.Fatal(app.Listen(address))
}
