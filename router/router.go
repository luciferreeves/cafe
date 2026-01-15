package router

import (
	"cafe/utils/urls"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Static("/static", "./static")

	urls.Attach(router)
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).SendString(err.Error())
}
