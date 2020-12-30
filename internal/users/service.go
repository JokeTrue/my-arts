package users

import "github.com/JokeTrue/my-arts/internal/users/domain"

// Service users interface for DB access
//go:generate mockgen -source=service.go -destination=service_mock.go -package=users Service
type Service interface {
	Delete(id int) error
	GetUsers() ([]domain.User, error)
	Create(user domain.User) (int, error)
	GetUserByID(id int) (*domain.User, error)
	Update(user domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
