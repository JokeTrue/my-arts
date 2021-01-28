package app

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JokeTrue/my-arts/internal/products"
	productsDelivery "github.com/JokeTrue/my-arts/internal/products/delivery/http"
	productsRepo "github.com/JokeTrue/my-arts/internal/products/repository/mysql"
	productsUseCase "github.com/JokeTrue/my-arts/internal/products/usecase"

	"github.com/JokeTrue/my-arts/internal/reviews"
	reviewsDelivery "github.com/JokeTrue/my-arts/internal/reviews/delivery/http"
	reviewsRepo "github.com/JokeTrue/my-arts/internal/reviews/repository/mysql"
	reviewsUseCase "github.com/JokeTrue/my-arts/internal/reviews/usecase"

	"github.com/JokeTrue/my-arts/internal/users"
	usersDelivery "github.com/JokeTrue/my-arts/internal/users/delivery/http"
	usersRepo "github.com/JokeTrue/my-arts/internal/users/repository/mysql"
	usersUseCase "github.com/JokeTrue/my-arts/internal/users/usecase"

	"github.com/JokeTrue/my-arts/internal/categories"
	categoriesDelivery "github.com/JokeTrue/my-arts/internal/categories/delivery/http"
	categoriesRepo "github.com/JokeTrue/my-arts/internal/categories/repository/mysql"
	categoriesUseCase "github.com/JokeTrue/my-arts/internal/categories/usecase"

	"github.com/JokeTrue/my-arts/internal/tags"
	tagsDelivery "github.com/JokeTrue/my-arts/internal/tags/delivery/http"
	tagsRepo "github.com/JokeTrue/my-arts/internal/tags/repository/mysql"
	tagsUseCase "github.com/JokeTrue/my-arts/internal/tags/usecase"

	"github.com/JokeTrue/my-arts/internal/friendship"
	friendshipDelivery "github.com/JokeTrue/my-arts/internal/friendship/delivery/http"
	friendshipRepo "github.com/JokeTrue/my-arts/internal/friendship/repository/mysql"
	friendshipUseCase "github.com/JokeTrue/my-arts/internal/friendship/usecase"

	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/JokeTrue/my-arts/pkg/logging"
	_ "github.com/JokeTrue/my-arts/pkg/tzinit" // Set TZ to UTC
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Settings struct {
	Debug           bool   `env:"DEBUG" envDefault:"false"`
	Port            string `env:"PORT" envDefault:"8080"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
	ShutdownTimeout int    `env:"SHUTDOWN_TIMEOUT" envDefault:"30"`
	SecretKey       string `env:"SECRET_KEY"`
	MigrationsPath  string `env:"MIGRATIONS_PATH"`
}

type Application struct {
	settings Settings
	db       *sqlx.DB
	router   *gin.Engine
	logger   logging.Logger

	usersUseCase      users.UseCase
	productsUseCase   products.UseCase
	reviewsUseCase    reviews.UseCase
	categoriesUseCase categories.UseCase
	tagsUseCase       tags.UseCase
	friendshipUseCase friendship.UseCase
}

func NewApplication(logger logging.Logger, settings Settings) *Application {
	router := gin.Default()
	router.Use(gin.Recovery())

	if gin.Mode() == "release" {
		router.Static("/static/", "./frontend/static")
		router.LoadHTMLGlob("./frontend/*.html")
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		router.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}

	application := &Application{
		settings: settings,
		logger:   logger,
		router:   router,
	}

	// 1. Setup Database
	application.setupDB()
	//application.setupMockData("db/mock_data/")

	// 2. Setup UseCases + Endpoints
	application.setupUseCases()

	// 3. Setup JWT Authentication
	apiGroup := application.setupJWT()

	// 4. Setup HTTP Endpoints
	usersDelivery.RegisterHTTPEndpoints(apiGroup, application.router, application.usersUseCase)
	productsDelivery.RegisterHTTPEndpoints(apiGroup, application.productsUseCase)
	reviewsDelivery.RegisterHTTPEndpoints(apiGroup, application.reviewsUseCase)
	categoriesDelivery.RegisterHTTPEndpoints(apiGroup, application.categoriesUseCase)
	tagsDelivery.RegisterHTTPEndpoints(apiGroup, application.tagsUseCase)
	friendshipDelivery.RegisterHTTPEndpoints(apiGroup, application.friendshipUseCase)

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

func (a *Application) setupUseCases() {
	// Users
	usersRepository := usersRepo.NewUsersRepository(a.db)
	a.usersUseCase = usersUseCase.NewUsersUseCase(usersRepository)

	// Products
	productsRepository := productsRepo.NewProductsRepository(a.db)
	a.productsUseCase = productsUseCase.NewProductsUseCase(productsRepository)

	// Reviews
	reviewsRepository := reviewsRepo.NewProductsRepository(a.db)
	a.reviewsUseCase = reviewsUseCase.NewReviewsUseCase(reviewsRepository)

	// Categories
	categoriesRepository := categoriesRepo.NewCategoriesRepository(a.db)
	a.categoriesUseCase = categoriesUseCase.NewCategoriesUseCase(categoriesRepository)

	// Tags
	tagsRepository := tagsRepo.NewTagsRepository(a.db)
	a.tagsUseCase = tagsUseCase.NewTagsUseCase(tagsRepository)

	// Friendship
	friendshipRepository := friendshipRepo.NewFriendshipRepository(a.db)
	a.friendshipUseCase = friendshipUseCase.NewFriendshipUseCase(friendshipRepository)
}

func (a *Application) setupJWT() *gin.RouterGroup {
	authMiddleware, err := jwt.GetJWTMiddleware(a.usersUseCase, a.settings.SecretKey)
	if err != nil {
		a.logger.WithError(err).Panic("failed to setup jwt")
	}
	a.router.POST("/api/login", authMiddleware.LoginHandler)

	apiGroup := a.router.Group("/api")
	apiGroup.Use(authMiddleware.MiddlewareFunc())
	apiGroup.GET("/auth/refresh_token", authMiddleware.RefreshHandler)

	return apiGroup
}

func (a *Application) setupDB() {
	db, err := sqlx.Connect("mysql", a.settings.DatabaseDSN+"?parseTime=true")
	if err != nil {
		a.logger.WithError(err).Panic("failed to setup db connection")
	}

	m, err := migrate.New("file://"+a.settings.MigrationsPath, "mysql://"+a.settings.DatabaseDSN)
	if err != nil {
		a.logger.WithError(err).Panic("failed to setup migrations")
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		a.logger.WithError(err).Panic("failed to up migrations")
	}

	a.db = db
}

func (a *Application) setupMockData(mockPath string) {
	if !a.settings.Debug {
		return
	}

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
			_ = tx.Rollback()
			return
		}

		if _, err := a.db.Exec(string(script)); err != nil {
			a.logger.WithError(err).Error("failed to load mock data")
			_ = tx.Rollback()
			return
		}
	}

	_ = tx.Commit()
}
