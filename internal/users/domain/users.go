package domain

import "time"

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email,omitempty" db:"email"`
	Password string `json:"password,omitempty" db:"password"`

	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Age       int    `json:"age" db:"age"`
	Gender    string `json:"gender" db:"gender"`
	Location  string `json:"location" db:"location"`
	Biography string `json:"biography" db:"biography"`

	Permissions string    `db:"permissions"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Service instance to manage Users
type Service struct {
	s Store
}

// NewService creates new User Service
func NewService(store Store) *Service {
	return &Service{s: store}
}

// Delete single User
func (svc Service) Delete(id int) error {
	return svc.s.delete(id)
}

// GetUserByID returns single User
func (svc Service) GetUserByID(id int) (*User, error) {
	return svc.s.getByID(id)
}

// GetUserByEmail returns single User
func (svc Service) GetUserByEmail(email string) (*User, error) {
	return svc.s.getByEmail(email)
}

// GetUsers returns all Users
func (svc Service) GetUsers() ([]User, error) {
	return svc.s.getAll()
}

// Create a User
func (svc Service) Create(user User) (int, error) {
	return svc.s.create(user)
}

// Update single User
func (svc Service) Update(user User) (*User, error) {
	return svc.s.update(user)
}
