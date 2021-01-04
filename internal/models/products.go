package models

import "time"

type Category struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Product struct {
	ID                 int       `json:"id" db:"id"`
	UserID             int       `json:"user_id" db:"user_id"`
	CategoryID         int       `json:"category_id" db:"category_id"`
	Title              string    `json:"title" db:"title"`
	GeneralDescription string    `json:"general_description" db:"general_description"`
	ProductDescription string    `json:"product_description" db:"product_description"`
	State              string    `json:"state" db:"state"`
	PriceAmount        float64   `json:"price_amount" db:"price_amount"`
	PriceCurrency      string    `json:"price_currency" db:"price_currency"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
}

type ProductTag struct {
	ID        int    `json:"id" db:"id"`
	ProductID int    `json:"product_id" db:"product_id"`
	Title     string `json:"title" db:"title"`
}

type ProductPhoto struct {
	ID        int       `json:"id" db:"id"`
	ProductID int       `json:"product_id" db:"product_id"`
	URL       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
