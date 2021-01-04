package products

import "github.com/JokeTrue/my-arts/internal/models"

type UseCase interface {
	Create(product models.Product) (int, error)
	GetProduct(id int) (*models.Product, error)
	Update(product models.Product) (int, error)
	Delete(id int) error

	GetProducts(states []string) ([]*models.Product, error)
	GetUserProducts(userId int, states []string) ([]*models.Product, error)
}
