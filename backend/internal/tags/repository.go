package tags

import "github.com/JokeTrue/my-arts/internal/models"

type Repository interface {
	Delete(id int) error
	GetTags() ([]*models.ProductTag, error)
	Create(tag models.ProductTag) (int, error)
	GetTag(id int) (*models.ProductTag, error)
	Update(tag models.ProductTag) (*models.ProductTag, error)
	GetProductTags(productId int) ([]*models.ProductTag, error)
}
