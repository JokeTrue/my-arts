package http

type UpdateReviewRequest struct {
	Comment             string `json:"comment" binding:"required"`
	CommunicationRating int    `json:"communication_rating" binding:"required"`
	DeliveryRating      int    `json:"delivery_rating" binding:"required"`
	AccuracyRating      int    `json:"accuracy_rating" binding:"required"`
}

type CreateReviewRequest struct {
	UpdateReviewRequest
	UserID int `json:"user_id" binding:"required"`
}
