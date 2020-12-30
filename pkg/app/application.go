package app

import (
	"github.com/JokeTrue/my-arts/internal/users"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JokeTrue/my-arts/pkg/logging"
)

type Application struct {
	debug  bool
	db     *sqlx.DB
	router *gin.Engine
	logger logging.Logger

	userRouter *users.Router
}

func NewApplication(debug bool, logger logging.Logger, dbDSN string) *Application {
	router := gin.Default()

	application := &Application{
		debug:  debug,
		logger: logger,
		router: router,
	}

	application.setupDB(dbDSN)
	if application.debug {
		application.setupMockData("db/mock_data/")
	}

	authMiddleware, err := jwt.GetJWTMiddleware(application.db)
	if err != nil {
		application.logger.WithError(err).Panic("failed to setup jwt")
	}

	// Auth
	router.POST("/login", authMiddleware.LoginHandler)

	apiGroup := router.Group("/api")
	apiGroup.Use(authMiddleware.MiddlewareFunc())
	{
		// Setup JWT Handlers
		apiGroup.GET("/auth/refresh_token", authMiddleware.RefreshHandler)

		// Setup Users Router
		users.NewRouter(application.db, logger).SetupRoutes(apiGroup)
	}

	return application
}

func (a *Application) Run() http.Handler {
	return a.router
}

func (a *Application) Stop() {
	if a.db == nil {
		return
	}
	if err := a.db.Close(); err != nil {
		a.logger.WithError(err).Error("failed to close database connection")
	}
}

func (a *Application) setupDB(databaseDSN string) {
	db, err := sqlx.Connect("mysql", databaseDSN)
	if err != nil {
		a.logger.WithError(err).Panic("failed to setup db connection")
	}

	m, err := migrate.New("file://db/migrations", "mysql://"+databaseDSN)
	if err != nil {
		a.logger.WithError(err).Panic("failed to setup migrations")
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		a.logger.WithError(err).Panic("failed to up migrations")
	}

	a.db = db
}

func (a *Application) setupMockData(mockPath string) {
	var files []string

	if err := filepath.Walk(mockPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	}); err != nil {
		panic(err)
	}

	tx, err := a.db.Beginx()
	if err != nil {
		a.logger.WithError(err).Error("failed to load mock data")
	}

	for _, file := range files {
		script, err := ioutil.ReadFile(file)
		if err != nil {
			a.logger.WithError(err).Error("failed to load mock data")
			tx.Rollback()
			return
		}

		if _, err := a.db.Exec(string(script)); err != nil {
			a.logger.WithError(err).Error("failed to load mock data")
			tx.Rollback()
			return
		}
	}

	_ = tx.Commit()
}
