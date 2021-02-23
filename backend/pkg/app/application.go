package app

import (
	"github.com/JokeTrue/my-arts/internal/products"
	productsDelivery "github.com/JokeTrue/my-arts/internal/products/delivery/http"
	productsRepo "github.com/JokeTrue/my-arts/internal/products/repository/mysql"
	productsUseCase "github.com/JokeTrue/my-arts/internal/products/usecase"
	"github.com/JokeTrue/my-arts/pkg/middleware"
	"github.com/JokeTrue/my-arts/pkg/utils"
	"net/http"

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
	Debug bool   `env:"DEBUG" envDefault:"false"`
	Port  string `env:"PORT" envDefault:"8080"`

	MasterDatabaseDSN string   `env:"MASTER_DB_DSN"`
	SlaveDatabaseDSNs []string `env:"SLAVE_DB_DSNS" envSeparator:","`

	ShutdownTimeout int    `env:"SHUTDOWN_TIMEOUT" envDefault:"30"`
	SecretKey       string `env:"SECRET_KEY"`
	MigrationsPath  string `env:"MIGRATIONS_PATH"`
}

type Application struct {
	settings Settings

	writeDB *sqlx.DB
	readDBs []*sqlx.DB

	router *gin.Engine
	logger logging.Logger

	usersUseCase      users.UseCase
	productsUseCase   products.UseCase
	reviewsUseCase    reviews.UseCase
	categoriesUseCase categories.UseCase
	tagsUseCase       tags.UseCase
	friendshipUseCase friendship.UseCase
}

func NewApplication(logger logging.Logger, settings Settings) *Application {
	router := gin.New()
	router.Use(middleware.RequestCancelRecover())

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
	application.setupDatabases()
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
	if a.writeDB == nil && a.readDBs == nil {
		return
	}

	if err := a.writeDB.Close(); err != nil {
		a.logger.WithError(err).Error("failed to close database connection")
	}

	for _, slave := range a.readDBs {
		if err := slave.Close(); err != nil {
			a.logger.WithError(err).Error("failed to close slave database connection")
		}
	}
}

func (a *Application) setupUseCases() {
	// Users
	usersRepository := usersRepo.NewUsersRepository(a.writeDB, utils.GetReadDatabase(a.readDBs))
	a.usersUseCase = usersUseCase.NewUsersUseCase(usersRepository)

	// Products
	productsRepository := productsRepo.NewProductsRepository(a.writeDB, utils.GetReadDatabase(a.readDBs))
	a.productsUseCase = productsUseCase.NewProductsUseCase(productsRepository)

	// Reviews
	reviewsRepository := reviewsRepo.NewProductsRepository(a.writeDB, utils.GetReadDatabase(a.readDBs))
	a.reviewsUseCase = reviewsUseCase.NewReviewsUseCase(reviewsRepository)

	// Categories
	categoriesRepository := categoriesRepo.NewCategoriesRepository(a.writeDB, utils.GetReadDatabase(a.readDBs))
	a.categoriesUseCase = categoriesUseCase.NewCategoriesUseCase(categoriesRepository)

	// Tags
	tagsRepository := tagsRepo.NewTagsRepository(a.writeDB, utils.GetReadDatabase(a.readDBs))
	a.tagsUseCase = tagsUseCase.NewTagsUseCase(tagsRepository)

	// Friendship
	friendshipRepository := friendshipRepo.NewFriendshipRepository(a.writeDB, utils.GetReadDatabase(a.readDBs))
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

func (a *Application) setupDatabases() {
	var err error

	// 1. Setup Master Database
	a.writeDB, err = sqlx.Connect("mysql", a.settings.MasterDatabaseDSN+"?parseTime=true")
	if err != nil {
		a.logger.WithError(err).Panic("failed to setup db connection")
	}
	a.writeDB.SetMaxOpenConns(300)

	m, err := migrate.New("file://"+a.settings.MigrationsPath, "mysql://"+a.settings.MasterDatabaseDSN)
	if err != nil {
		a.logger.WithError(err).Panic("failed to setup migrations")
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		a.logger.WithError(err).Panic("failed to up migrations")
	}

	// 2. Setup Slave Databases
	if a.settings.SlaveDatabaseDSNs == nil {
		a.logger.Panic("empty slave databases list")
	}

	for _, dsn := range a.settings.SlaveDatabaseDSNs {
		slave, err := sqlx.Connect("mysql", dsn+"?parseTime=true")
		if err != nil {
			a.logger.WithError(err).Panic("failed to setup slave db connection")
		}
		slave.SetMaxOpenConns(300)
		a.readDBs = append(a.readDBs, slave)
	}
}
