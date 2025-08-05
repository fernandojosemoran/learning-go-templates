package main

import (
	"context"
	"time"

	blog "github.com/fernandojosemoran/go-templates/internal/apps/blog/application/migrations"
	dbConf "github.com/fernandojosemoran/go-templates/internal/apps/blog/infrastructure/db"
	"github.com/fernandojosemoran/go-templates/internal/middlewares"
	"github.com/fernandojosemoran/go-templates/internal/routes"
	config "github.com/fernandojosemoran/go-templates/pkg/config/constant"
	"github.com/fernandojosemoran/go-templates/pkg/logger"
	"github.com/fernandojosemoran/go-templates/pkg/server"
	"gorm.io/gorm"
)

func main() {
	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Second*2)

	err := dbConf.InitializationDB(&dbConf.DatabaseOptions{
		Context: ctx,
		Migrations: []func(db *gorm.DB) error{
			blog.BlogMigrations,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	} else {
		logger.Info("⭐ Postgres database connected.")
	}

	defer ctxCancel()

	var srv = server.HttpServer{
		Port:        config.Env().PORT,
		Middlewares: middlewares.Middlewares,
		Handlers:    routes.Handlers,
	}

	srv.Start()
}
