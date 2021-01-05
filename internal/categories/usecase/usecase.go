package usecase

import (
	"github.com/JokeTrue/my-arts/internal/categories"
	"github.com/JokeTrue/my-arts/internal/models"
)

type CategoriesUseCase struct {
	repo categories.Repository
}

func NewCategoriesUseCase(repo categories.Repository) *CategoriesUseCase {
	return &CategoriesUseCase{repo: repo}
}

func (u *CategoriesUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *CategoriesUseCase) GetCategories() ([]*models.Category, error) {
	return u.repo.GetCategories()
}

func (u *CategoriesUseCase) GetCategory(id int) (*models.Category, error) {
	return u.repo.GetCategory(id)
}

func (u *CategoriesUseCase) Create(category models.Category) (int, error) {
	return u.repo.Create(category)
}

func (u *CategoriesUseCase) Update(category models.Category) (*models.Category, error) {
	return u.repo.Update(category)
}
