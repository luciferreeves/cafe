package router

import (
	"cafe/types"
	"cafe/utils/shortcuts"
	"cafe/utils/urls"

	"github.com/gofiber/fiber/v2"
)

func init() {
	urls.SetNamespace("")

	urls.Path(types.GET, "/", func(c *fiber.Ctx) error {
		return shortcuts.Render(c, "pages/home", nil)
	}, "home")
}
