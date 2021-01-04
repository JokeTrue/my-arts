package http

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/reviews"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckObjectPermissions(c *gin.Context, useCase reviews.UseCase) (*models.Review, error) {
	reviewerId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		return nil, err
	}

	rawId := c.Param("review_id")
	reviewId, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, err
	}

	review, err := useCase.GetReview(reviewId)
	if err != nil {
		return nil, err
	}

	if review.ReviewerID != reviewerId {
		return nil, reviews.ErrReviewPermissionDenied
	}

	return review, nil
}
