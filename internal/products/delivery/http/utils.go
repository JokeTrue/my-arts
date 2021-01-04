package http

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/products"
	"github.com/JokeTrue/my-arts/pkg/utils"
)

func ValidateProductState(states []string) error {
	for _, state := range states {
		if !utils.Contains(models.AllProductStates, state) {
			return products.ErrUnknownProductState
		}
	}
	return nil
}
