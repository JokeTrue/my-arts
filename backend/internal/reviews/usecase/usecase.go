package usecase

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/reviews"
)

type ReviewsUseCase struct {
	repo reviews.Repository
}

func NewReviewsUseCase(repo reviews.Repository) *ReviewsUseCase {
	return &ReviewsUseCase{repo: repo}
}

func (u *ReviewsUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *ReviewsUseCase) Create(review models.Review) (int, error) {
	return u.repo.Create(review)
}

func (u *ReviewsUseCase) Update(review models.Review) (*models.Review, error) {
	return u.repo.Update(review)
}

func (u *ReviewsUseCase) GetUserReviews(userId int) ([]*models.Review, error) {
	return u.repo.GetUserReviews(userId)
}

func (u *ReviewsUseCase) GetReview(id int) (*models.Review, error) {
	return u.repo.GetReview(id)
}
