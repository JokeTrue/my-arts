package models

import (
	"time"

	"github.com/JokeTrue/my-arts/pkg/utils"
)

type ReviewUser struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

type Review struct {
	ID                  int                    `json:"id" db:"id"`
	UserID              int                    `json:"-" db:"user_id"`
	ReviewerID          int                    `json:"-" db:"reviewer_id"`
	Reviewer            ReviewUser             `json:"reviewer" db:"reviewer"`
	Comment             string                 `json:"comment" db:"comment"`
	CommunicationRating int                    `json:"communication_rating" db:"communication_rating"`
	DeliveryRating      int                    `json:"delivery_rating" db:"delivery_rating"`
	AccuracyRating      int                    `json:"accuracy_rating" db:"accuracy_rating"`
	EditedAt            utils.NullTimeWithJSON `json:"edited_at" db:"edited_at"`
	CreatedAt           time.Time              `json:"created_at" db:"created_at"`
}
