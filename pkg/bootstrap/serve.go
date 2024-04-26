package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
