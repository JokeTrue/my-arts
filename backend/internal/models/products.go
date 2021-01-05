package models

import "time"

const (
	StateNew        = "NEW"
	StateInProgress = "IN_PROGRESS"
	StateSold       = "SOLD"
	StateClosed     = "CLOSED"
	StateArchived   = "ARCHIVED"
)

var (
	AllProductStates = []string{
		StateNew,
		StateInProgress,
		StateSold,
		StateClosed,
		StateArchived,
	}
	AllowedForSearchStates = []string{
		StateSold,
		StateInProgress,
	}
	AllowedForUserSearchStates = []string{
		StateNew,
		StateInProgress,
		StateSold,
		StateArchived,
	}
)

type Product struct {
	ID                 int       `json:"id" db:"id"`
	UserID             int       `json:"user_id" db:"user_id"`
	CategoryID         int       `json:"-" db:"category_id"`
	Title              string    `json:"title" db:"title"`
	GeneralDescription string    `json:"general_description" db:"general_description"`
	ProductDescription string    `json:"product_description" db:"product_description"`
	State              string    `json:"state" db:"state"`
	PriceAmount        float64   `json:"price_amount" db:"price_amount"`
	PriceCurrency      string    `json:"price_currency" db:"price_currency"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`

	// Foreign Keys Objects
	User     *User           `json:"-" db:"user"`
	Category *Category       `json:"category" db:"category"`
	Tags     []*ProductTag   `json:"tags" db:"-"`
	Photos   []*ProductPhoto `json:"photos" db:"-"`
}

type Category struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ProductTag struct {
	ID        int    `json:"id" db:"id"`
	ProductID int    `json:"-" db:"product_id"`
	Title     string `json:"title" db:"title"`
}

type ProductPhoto struct {
	ID        int       `json:"id" db:"id"`
	ProductID int       `json:"-" db:"product_id"`
	URL       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}
