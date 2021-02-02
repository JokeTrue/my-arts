package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/users"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase users.UseCase
}

func NewHandler(useCase users.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var request SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if request.Password1 != request.Password2 {
		appErrors.JSONError(c, users.ErrUserPasswordsDontMatch, nil)
		return
	}

	existingUser, err := h.useCase.GetUserByEmail(request.Email)
	if err != nil && err != users.ErrUserNotFound {
		appErrors.JSONError(c, err, nil)
		return
	}
	if existingUser != nil {
		appErrors.JSONError(c, users.ErrUserAlreadyExists, nil)
		return
	}

	user := models.User{
		Email:     request.Email,
		Password:  request.Password1,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Age:       request.Age,
		Gender:    request.Gender,
		Location:  request.Location,
		Biography: request.Biography,
		CreatedAt: time.Now(),
	}

	user.ID, err = h.useCase.Create(user)
	if err != nil {
		appErrors.JSONError(c, err, user)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetCurrentUser(c *gin.Context) {
	userId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	user, err := h.useCase.GetUserByID(userId)
	if err != nil {
		appErrors.JSONError(c, err, user)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUser(c *gin.Context) {
	rawUserId := c.Param("user_id")
	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		appErrors.JSONError(c, err, rawUserId)
		return
	}

	user, err := h.useCase.GetUserByID(userId)
	if err != nil {
		appErrors.JSONError(c, err, user)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) SearchUsers(c *gin.Context) {
	query := c.Query("query")

	usersList, err := h.useCase.SearchUsers(query)
	if err != nil {
		appErrors.JSONError(c, err, query)
		return
	}

	c.JSON(http.StatusOK, usersList)
}

func (h *Handler) GetUserFriends(c *gin.Context) {
	userId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	usersList, err := h.useCase.GetUserFriends(userId)
	if err != nil {
		appErrors.JSONError(c, err, userId)
		return
	}

	c.JSON(http.StatusOK, usersList)
}
