package main

import (
	"flag"
	"time"

	"github.com/JokeTrue/my-arts/pkg/middleware"
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

var (
	addr            string
	databaseDSN     string
	debug           bool
	shutdownTimeout time.Duration
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Debug Mode")
	flag.StringVar(&addr, "addr", ":9080", "App addr")
	flag.StringVar(&databaseDSN, "db-dsn", "", "Database DSN")
	flag.DurationVar(&shutdownTimeout, "shutdown-timeout", 30*time.Second, "Graceful shutdown timeout")
}

func main() {
	flag.Parse()

	// 1. Setup Application
	logger := logging.DefaultLogger
	application := app.NewApplication(debug, logger, databaseDSN)

	// 2. Setup HTTP Server
	srv := service.NewHTTPServer(addr, shutdownTimeout, alice.New(
		gziphandler.GzipHandler,
		middleware.Logger(logger),
	).Then(application.Run()))

	// 3. Run Service
	service.Run(srv, appName)

	// 4. Stop Application
	application.Stop()
}
