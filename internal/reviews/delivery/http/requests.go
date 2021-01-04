package http

type UpdateReviewRequest struct {
	Comment             string `json:"comment"`
	CommunicationRating int    `json:"communication_rating"`
	DeliveryRating      int    `json:"delivery_rating"`
	AccuracyRating      int    `json:"accuracy_rating"`
}

type CreateReviewRequest struct {
	UpdateReviewRequest
	UserID int `json:"user_id"`
}
