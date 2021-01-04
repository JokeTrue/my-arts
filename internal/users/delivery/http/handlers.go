package http

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/users"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase users.UseCase
}

func NewHandler(useCase users.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userId, err := h.useCase.Create(user)
	if err != nil {
		appErrors.JSONError(c, err, user)
		return
	}

	c.JSON(http.StatusCreated, map[string]int{"user_id": userId})
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
