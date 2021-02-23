package mysql

import (
	"database/sql"
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/reviews"
	"github.com/JokeTrue/my-arts/internal/users"
	"github.com/jmoiron/sqlx"
)

type ReviewsRepository struct {
	writeDB *sqlx.DB
	readDB  *sqlx.DB
}

func NewProductsRepository(writeDB, readDB *sqlx.DB) *ReviewsRepository {
	return &ReviewsRepository{writeDB: writeDB, readDB: readDB}
}

func (r *ReviewsRepository) GetReview(id int) (*models.Review, error) {
	var review models.Review
	if err := r.readDB.Get(&review, QueryGetReviewByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, reviews.ErrReviewNotFound
		}
		return nil, reviews.ErrReviewQuery
	}

	return &review, nil
}

func (r *ReviewsRepository) Delete(id int) error {
	res, err := r.writeDB.Exec(QueryDeleteReview, id)
	if err != nil {
		return users.ErrUserQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return users.ErrUserQuery
	}
	if affect == 0 {
		return users.ErrUserNotFound
	}

	return nil
}

func (r *ReviewsRepository) Create(review models.Review) (int, error) {
	result, err := r.writeDB.Exec(
		QueryCreateReview,
		review.UserID,
		review.ReviewerID,
		review.Comment,
		review.DeliveryRating,
		review.CommunicationRating,
		review.AccuracyRating,
	)
	if err != nil {
		return 0, reviews.ErrReviewQuery
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, reviews.ErrReviewQuery
	}

	return int(lastID), nil
}

func (r *ReviewsRepository) Update(review models.Review) (*models.Review, error) {
	res, err := r.writeDB.Exec(
		QueryUpdateReview,
		review.Comment,
		review.DeliveryRating,
		review.CommunicationRating,
		review.AccuracyRating,
		review.ID,
	)
	if err != nil {
		return nil, reviews.ErrReviewQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return nil, reviews.ErrReviewQuery
	}
	if affect == 0 {
		return nil, reviews.ErrReviewNotFound
	}

	return &review, nil
}

func (r *ReviewsRepository) GetUserReviews(userId int) ([]*models.Review, error) {
	list := []*models.Review{}
	if err := r.readDB.Select(&list, QueryGetUserReviews, userId); err != nil {
		return nil, users.ErrUserQuery
	}
	return list, nil
}
