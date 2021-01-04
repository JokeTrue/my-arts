package http

import (
	"net/http"
	"strconv"

	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/reviews"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase reviews.UseCase
}

func NewHandler(useCase reviews.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) Delete(c *gin.Context) {
	rawReviewId := c.Param("review_id")
	reviewId, err := strconv.Atoi(rawReviewId)
	if err != nil {
		appErrors.JSONError(c, err, rawReviewId)
		return
	}

	if _, err = CheckObjectPermissions(c, h.useCase); err != nil {
		appErrors.JSONError(c, err, reviewId)
		return
	}

	err = h.useCase.Delete(reviewId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) Create(c *gin.Context) {
	var request CreateReviewRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	reviewerId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	if request.UserID == reviewerId {
		appErrors.JSONError(c, reviews.ErrReviewPermissionDenied, nil)
		return
	}

	review := models.Review{
		ReviewerID:          reviewerId,
		UserID:              request.UserID,
		Comment:             request.Comment,
		DeliveryRating:      request.DeliveryRating,
		AccuracyRating:      request.AccuracyRating,
		CommunicationRating: request.CommunicationRating,
	}

	review.ID, err = h.useCase.Create(review)
	if err != nil {
		appErrors.JSONError(c, err, review)
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (h *Handler) Update(c *gin.Context) {
	var request UpdateReviewRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	review, err := CheckObjectPermissions(c, h.useCase)
	if err != nil {
		appErrors.JSONError(c, err, review)
		return
	}

	review.Comment = request.Comment
	review.DeliveryRating = request.DeliveryRating
	review.AccuracyRating = request.AccuracyRating
	review.CommunicationRating = request.AccuracyRating

	updatedReview, err := h.useCase.Update(*review)
	if err != nil {
		appErrors.JSONError(c, err, review)
		return
	}

	c.JSON(http.StatusOK, updatedReview)
}

func (h *Handler) GetUserReviews(c *gin.Context) {
	rawUserId := c.Param("user_id")
	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		appErrors.JSONError(c, err, userId)
		return
	}

	userReviews, err := h.useCase.GetUserReviews(userId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, userReviews)
}

func (h *Handler) GetReview(c *gin.Context) {
	rawReviewId := c.Param("review_id")
	reviewId, err := strconv.Atoi(rawReviewId)
	if err != nil {
		appErrors.JSONError(c, err, rawReviewId)
		return
	}

	review, err := h.useCase.GetReview(reviewId)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, review)
}
