package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/database"
	"blog/pkg/html"
	"blog/pkg/routing"
	"blog/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHtml(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
