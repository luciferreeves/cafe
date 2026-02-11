package processors

import (
	"cafe/repositories"
	"cafe/session"

	"github.com/gofiber/fiber/v2"
)

func user(ctx *fiber.Ctx) error {
	sess, err := session.Store.Get(ctx)
	if err != nil {
		return ctx.Next()
	}

	username := sess.Get("username")
	if username == nil {
		return ctx.Next()
	}

	user, err := repositories.GetUserByUsername(username.(string))
	if err != nil {
		return ctx.Next()
	}

	ctx.Locals("User", user)

	return ctx.Next()
}
