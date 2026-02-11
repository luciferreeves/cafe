package controllers

import (
	"cafe/utils/meta"
	"cafe/utils/shortcuts"

	"github.com/gofiber/fiber/v2"
)

func MainHall(context *fiber.Ctx) error {
	meta.SetPageTitle(context, "Main Hall")

	return shortcuts.Render(context, "pages/main", nil)
}
