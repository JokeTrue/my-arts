package users

import (
	"github.com/JokeTrue/my-arts/internal/models"
)

type Repository interface {
	Delete(id int) error
	GetUsers() ([]models.User, error)
	Create(user models.User) (int, error)
	GetUserByID(id int) (*models.User, error)
	Update(user models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
