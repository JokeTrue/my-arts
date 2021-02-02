package models

import "time"

type shortUser struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Age       int    `json:"age" db:"age"`
	Location  string `json:"location" db:"location"`
}

type FriendshipRequest struct {
	ID        int       `json:"id" db:"id"`
	ActorID   int       `json:"-" db:"actor_id"`
	FriendID  int       `json:"-" db:"friend_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Foreign Keys Objects
	User *shortUser `json:"user" db:"user"`
}
