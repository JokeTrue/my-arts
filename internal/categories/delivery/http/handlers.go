package http

import (
	"github.com/JokeTrue/my-arts/internal/categories"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase categories.UseCase
}

func NewHandler(useCase categories.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetCategories(c *gin.Context) {
	categoriesList, err := h.useCase.GetCategories()
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, categoriesList)
}
