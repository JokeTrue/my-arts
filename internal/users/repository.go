package users

import (
	"github.com/JokeTrue/my-arts/internal/models"
)

type Repository interface {
	Create(user models.User) (int, error)              // C
	GetUserByID(id int) (*models.User, error)          // R
	GetUserByEmail(email string) (*models.User, error) // R
	Update(user models.User) (*models.User, error)     // U
	Delete(id int) error                               // D

	GetUsers() ([]models.User, error)
}
