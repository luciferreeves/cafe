package controllers

import (
	"cafe/utils/auth"
	"cafe/utils/meta"
	"cafe/utils/shortcuts"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(context *fiber.Ctx) error {
	if auth.IsAuthenticated(context) {
		return shortcuts.Redirect(context, "mainHall")
	}

	meta.SetPageTitle(context, "Open ID Authentication")

	return shortcuts.Render(context, "pages/auth", nil)
}
