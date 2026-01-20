package meta

import (
	"cafe/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetPageTitle(context *fiber.Ctx, title string) {
	title = fmt.Sprintf("%s | %s", title, config.Server.AppName)
	context.Locals("Title", title)
}
