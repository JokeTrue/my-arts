package products

import "github.com/JokeTrue/my-arts/internal/models"

type Repository interface {
	Create(product models.Product) (int, error) // C
	GetProduct(id int) (*models.Product, error) // R
	Update(product models.Product) (int, error) // U
	Delete(id int) error                        // D

	GetProducts(states []string) ([]*models.Product, error)
	GetUserProducts(userId int, states []string) ([]*models.Product, error)
}
