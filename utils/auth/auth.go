package auth

import (
	"cafe/session"
	"cafe/utils/shortcuts"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(context *fiber.Ctx) bool {
	session, err := session.Store.Get(context)
	if err != nil {
		return false
	}

	username := session.Get("username")
	return username != nil
}

func RequireAuthentication(handler fiber.Handler) fiber.Handler {
	return func(context *fiber.Ctx) error {
		if !IsAuthenticated(context) {
			return shortcuts.Redirect(context, "auth.authenticate")
		}
		return handler(context)
	}
}
