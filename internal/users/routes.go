package users

import (
	"github.com/JokeTrue/my-arts/internal/users/domain"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/JokeTrue/my-arts/pkg/logging"
	"github.com/JokeTrue/my-arts/pkg/middleware"
	"github.com/JokeTrue/my-arts/pkg/permissions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

// Router structs represents Users Handlers
type Router struct {
	service Service
	logger  logging.Logger
}

// NewRouter is creating NewStore User Router Handlers
func NewRouter(db *sqlx.DB, logger logging.Logger) *Router {
	store := domain.NewStore(db)
	return &Router{
		logger:  logger,
		service: domain.NewService(store),
	}
}

func (h *Router) SetupRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("", middleware.CheckPermissionsMiddleware(h.getUser, permissions.AdminPermission))
		users.POST("", h.createUser)
		users.GET("/list", h.getUsers)
		users.PUT("/:user_id", h.updateUser)
		users.DELETE("/:user_id", h.deleteUserByID)
	}
}

func (h *Router) getUser(c *gin.Context) {
	var (
		err  error
		user *domain.User
	)

	q := c.Request.URL.Query()
	email := q.Get("email")
	rawUserId := q.Get("user_id")

	if email != "" {
		user, err = h.service.GetUserByEmail(email)
	} else {
		userId, err := strconv.Atoi(rawUserId)
		if err != nil {
			appErrors.JSONError(c, appErrors.ErrBadParameter, "user_id")
			return
		}
		user, err = h.service.GetUserByID(userId)
	}
	if err != nil {
		appErrors.JSONError(c, err, map[string]string{"user_id": rawUserId, "email": email})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Router) getUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Router) deleteUserByID(c *gin.Context) {
	userId := c.Param("user_id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		appErrors.JSONError(c, appErrors.ErrBadParameter, "user_id")
		return
	}

	if err = h.service.Delete(id); err != nil {
		appErrors.JSONError(c, err, "user_id")
		return
	}
}

func (h *Router) updateUser(c *gin.Context) {
	userId := c.Param("user_id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		appErrors.JSONError(c, appErrors.ErrBadParameter, "user_id")
		return
	}

	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updatedUser, err := h.service.Update(
		domain.User{
			ID:        id,
			Age:       user.Age,
			Gender:    user.Gender,
			Location:  user.Location,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Biography: user.Biography,
		},
	)
	if err != nil {
		appErrors.JSONError(c, err, "user_id")
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *Router) createUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userId, err := h.service.Create(user)
	if err != nil {
		appErrors.JSONError(c, err, user)
		return
	}

	c.JSON(http.StatusCreated, map[string]int{"user_id": userId})
}
