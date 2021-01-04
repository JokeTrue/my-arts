package http

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/products"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/JokeTrue/my-arts/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ValidateProductStates(states []string, allowedStates []string) ([]string, error) {
	validatedStates := make([]string, 0, len(states))
	for _, state := range states {
		if !utils.Contains(models.AllProductStates, state) {
			return nil, products.ErrUnknownProductState
		}
		if utils.Contains(allowedStates, state) {
			validatedStates = append(validatedStates, state)
		}
	}
	return validatedStates, nil
}

func CheckObjectPermissions(c *gin.Context, useCase products.UseCase) (*models.Product, error) {
	userId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		return nil, err
	}

	rawProductId := c.Param("product_id")
	productId, err := strconv.Atoi(rawProductId)
	if err != nil {
		return nil, err
	}

	product, err := useCase.GetProduct(productId)
	if err != nil {
		return nil, err
	}

	if product.UserID == userId {

	}

	//TODO Permissions CHECKS

	return product, nil
}
