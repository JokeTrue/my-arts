package categories

import "github.com/JokeTrue/my-arts/internal/models"

type UseCase interface {
	Delete(id int) error
	GetCategories() ([]*models.Category, error)
	GetCategory(id int) (*models.Category, error)
	Create(category models.Category) (int, error)
	Update(category models.Category) (*models.Category, error)
}
