package users

import (
	"github.com/JokeTrue/my-arts/internal/models"
)

type Repository interface {
	Delete(id int) error
	GetUsers(offset, limit int) ([]models.User, error)
	Create(user models.User) (int, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	Update(user models.User) (*models.User, error)
	SearchUsers(query string, offset, limit int) ([]*models.User, error)
	GetUserFriends(id int, offset, limit int) ([]*models.User, error)
	GetTotalCount() (int, error)
}
