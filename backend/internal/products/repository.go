package products

import "github.com/JokeTrue/my-arts/internal/models"

type Repository interface {
	Delete(id int) error
	Create(product models.Product) (int, error)
	GetProduct(id int) (*models.Product, error)
	Update(product models.Product) (*models.Product, error)
	GetProducts(states []string) ([]*models.Product, error)
	GetUserProducts(userId int, states []string) ([]*models.Product, error)
}
