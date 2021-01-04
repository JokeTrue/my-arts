package usecase

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/products"
)

type ProductsUseCase struct {
	repo products.Repository
}

func NewProductsUseCase(repo products.Repository) *ProductsUseCase {
	return &ProductsUseCase{repo: repo}
}

func (u *ProductsUseCase) Create(product models.Product) (int, error) {
	return u.repo.Create(product)
}

func (u *ProductsUseCase) GetProduct(id int) (*models.Product, error) {
	return u.repo.GetProduct(id)
}

func (u *ProductsUseCase) Update(product models.Product) (int, error) {
	return u.repo.Update(product)
}

func (u *ProductsUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *ProductsUseCase) GetProducts(states []string) ([]*models.Product, error) {
	return u.repo.GetProducts(states)
}

func (u *ProductsUseCase) GetUserProducts(userId int, states []string) ([]*models.Product, error) {
	return u.repo.GetUserProducts(userId, states)
}
