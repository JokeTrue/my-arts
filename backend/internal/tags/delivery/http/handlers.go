package http

import (
	"github.com/JokeTrue/my-arts/internal/tags"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase tags.UseCase
}

func NewHandler(useCase tags.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetTags(c *gin.Context) {
	tagsList, err := h.useCase.GetTags()
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, tagsList)
}

func (h *Handler) GetProductTags(c *gin.Context) {
	rawProductId := c.Param("product_id")
	productId, err := strconv.Atoi(rawProductId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	tagsList, err := h.useCase.GetProductTags(productId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, tagsList)
}
