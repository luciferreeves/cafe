package router

import (
	"cafe/controllers"
	"cafe/types"
	"cafe/utils/urls"
)

func init() {
	urls.SetNamespace("auth")

	urls.Path(types.GET, "/login", controllers.Login, "login")
	urls.Path(types.GET, "/callback", controllers.Callback, "callback")
	urls.Path(types.GET, "/logout", controllers.Logout, "logout")
}
