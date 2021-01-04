package reviews

import "github.com/JokeTrue/my-arts/internal/models"

type Repository interface {
	Delete(id int) error
	GetReview(id int) (*models.Review, error)
	Create(review models.Review) (int, error)
	Update(review models.Review) (*models.Review, error)
	GetUserReviews(userId int) ([]*models.Review, error)
}
