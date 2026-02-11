package router

import (
	"cafe/controllers"
	"cafe/types"
	"cafe/utils/urls"
)

func init() {
	urls.SetNamespace("auth")

	urls.Path(types.GET, "/", controllers.Authenticate, "authenticate")
}
