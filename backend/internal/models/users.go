package models

import "time"

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email,omitempty" db:"email"`
	Password string `json:"-" db:"password"`

	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Age       int    `json:"age" db:"age"`
	Gender    string `json:"gender" db:"gender"`
	Location  string `json:"location" db:"location"`
	Biography string `json:"biography" db:"biography"`

	Permissions string    `json:"-" db:"permissions"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
