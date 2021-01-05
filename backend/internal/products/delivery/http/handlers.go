package http

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/products"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/JokeTrue/my-arts/pkg/utils"
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

	userId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	product, err := h.useCase.GetProduct(productId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	if userId != product.UserID && !utils.Contains(models.AllowedForSearchStates, product.State) {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	product, err := CheckObjectPermissions(c, h.useCase)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	err = h.useCase.Delete(product.ID)
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
	states, err = ValidateProductStates(states, models.AllowedForSearchStates)
	if err != nil {
		appErrors.JSONError(c, err, states)
		return
	}

	productsList, err = h.useCase.GetProducts(states)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, productsList)
}

func (h *Handler) GetUserProducts(c *gin.Context) {
	var (
		err    error
		states []string
	)

	if rawStates := c.Query("states"); rawStates != "" {
		states = strings.Split(rawStates, ",")
	}
	states, err = ValidateProductStates(states, models.AllowedForUserSearchStates)
	if err != nil {
		appErrors.JSONError(c, err, states)
		return
	}

	rawUserId := c.Param("user_id")
	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		appErrors.JSONError(c, err, rawUserId)
		return
	}

	productsList, err := h.useCase.GetUserProducts(userId, states)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, productsList)
}

func (h *Handler) NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, nil)
}
