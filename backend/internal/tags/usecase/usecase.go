package usecase

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/tags"
)

type TagsUseCase struct {
	repo tags.Repository
}

func NewTagsUseCase(repo tags.Repository) *TagsUseCase {
	return &TagsUseCase{repo: repo}
}

func (u *TagsUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *TagsUseCase) GetTags() ([]*models.ProductTag, error) {
	return u.repo.GetTags()
}

func (u *TagsUseCase) Create(tag models.ProductTag) (int, error) {
	return u.repo.Create(tag)
}

func (u *TagsUseCase) GetTag(id int) (*models.ProductTag, error) {
	return u.repo.GetTag(id)
}

func (u *TagsUseCase) Update(tag models.ProductTag) (*models.ProductTag, error) {
	return u.repo.Update(tag)
}

func (u *TagsUseCase) GetProductTags(productId int) ([]*models.ProductTag, error) {
	return u.repo.GetProductTags(productId)
}
