package processors

import (
	"cafe/config"

	"github.com/gofiber/fiber/v2"
)

const defaultTitle = "Shifoo's Cafe"

func metadata(ctx *fiber.Ctx) error {
	ctx.Locals("Title", defaultTitle)
	ctx.Locals("AppName", config.Server.AppName)
	ctx.Locals("AppDescription", config.Server.AppDescription)

	return ctx.Next()
}
