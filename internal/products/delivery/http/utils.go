package http

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/products"
	"github.com/JokeTrue/my-arts/pkg/utils"
)

func ValidateProductStates(states []string) ([]string, error) {
	validatedStates := make([]string, 0, len(states))
	for _, state := range states {
		if !utils.Contains(models.AllProductStates, state) {
			return nil, products.ErrUnknownProductState
		}
		if utils.Contains(models.AllowedForSearchStates, state) {
			validatedStates = append(validatedStates, state)
		}
	}
	return validatedStates, nil
}
