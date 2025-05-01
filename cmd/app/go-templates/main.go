package main

import (
	"github.com/fernandojosemoran/go-templates/internal/middlewares"
	"github.com/fernandojosemoran/go-templates/internal/routes"
	"github.com/fernandojosemoran/go-templates/pkg/server"
)

func main() {
	var srv = server.HttpServer{
		Port:        3000,
		Middlewares: middlewares.Middlewares,
		Handlers:    routes.Handlers,
	}

	srv.Start()
}
