package main

import (
	"time"

	"github.com/JokeTrue/my-arts/pkg/middleware"
	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"

	"github.com/JokeTrue/my-arts/pkg/app"

	"github.com/JokeTrue/my-arts/pkg/service"
	"github.com/NYTimes/gziphandler"
	"github.com/justinas/alice"

	"github.com/JokeTrue/my-arts/pkg/logging"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var appName = "my-arts-backend"

func main() {
	// 1. Setup Application
	var settings app.Settings
	if err := env.Parse(&settings); err != nil {
		logging.WithError(err).Fatal("failed to parse config")
	}

	logger := logging.DefaultLogger
	application := app.NewApplication(logger, settings)

	// 2. Setup HTTP Server
	srv := service.NewHTTPServer(":"+settings.Port, time.Duration(settings.ShutdownTimeout)*time.Second, alice.New(
		gziphandler.GzipHandler,
		middleware.Logger(logger),
	).Then(application.Run()))

	// 3. Run Service
	service.Run(srv, appName)

	// 4. Stop Application
	application.Stop()
}
