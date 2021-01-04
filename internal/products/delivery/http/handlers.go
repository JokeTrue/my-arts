package http

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/products"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	useCase products.UseCase
}

func NewHandler(useCase products.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetProduct(c *gin.Context) {
	rawProductId := c.Param("product_id")
	productId, err := strconv.Atoi(rawProductId)
	if err != nil {
		appErrors.JSONError(c, err, rawProductId)
		return
	}

	product, err := h.useCase.GetProduct(productId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	rawProductId := c.Param("product_id")
	productId, err := strconv.Atoi(rawProductId)
	if err != nil {
		appErrors.JSONError(c, err, rawProductId)
		return
	}

	err = h.useCase.Delete(productId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) GetProducts(c *gin.Context) {
	var (
		err          error
		states       []string
		productsList []*models.Product
	)
	if rawStates := c.Query("states"); rawStates != "" {
		states = strings.Split(rawStates, ",")
	}
	if err = ValidateProductState(states); err != nil {
		appErrors.JSONError(c, err, states)
		return
	}

	rawUserId := c.Query("user_id")
	if rawUserId != "" {
		userId, err := strconv.Atoi(rawUserId)
		if err != nil {
			appErrors.JSONError(c, err, nil)
			return
		}
		productsList, err = h.useCase.GetUserProducts(userId, states)
		if err != nil {
			appErrors.JSONError(c, err, nil)
			return
		}
	}

	if rawUserId == "" {
		productsList, err = h.useCase.GetProducts(states)
		if err != nil {
			appErrors.JSONError(c, err, nil)
			return
		}
	}

	c.JSON(http.StatusOK, productsList)
}

func (h *Handler) NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, nil)
}
