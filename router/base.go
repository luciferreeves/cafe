package router

import (
	"cafe/controllers"
	"cafe/types"
	"cafe/utils/auth"
	"cafe/utils/urls"
)

func init() {
	urls.SetNamespace("")

	urls.Path(types.GET, "/", auth.RequireAuthentication(controllers.MainHall), "mainHall")
}
