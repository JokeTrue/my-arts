package usecase

import (
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/users"
)

type UsersUseCase struct {
	repo users.Repository
}

func NewUsersUseCase(repo users.Repository) *UsersUseCase {
	return &UsersUseCase{repo: repo}
}

func (u *UsersUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *UsersUseCase) GetUsers(offset, limit int) ([]models.User, error) {
	return u.repo.GetUsers(offset, limit)
}

func (u *UsersUseCase) Create(user models.User) (int, error) {
	return u.repo.Create(user)
}

func (u *UsersUseCase) GetUserByID(id int) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UsersUseCase) Update(user models.User) (*models.User, error) {
	return u.repo.Update(user)
}

func (u *UsersUseCase) GetUserByEmail(email string) (*models.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *UsersUseCase) SearchUsers(query string, offset, limit int) ([]*models.User, error) {
	return u.repo.SearchUsers(query, offset, limit)
}

func (u *UsersUseCase) GetUserFriends(id int, offset, limit int) ([]*models.User, error) {
	return u.repo.GetUserFriends(id, offset, limit)
}

func (u *UsersUseCase) GetTotalCount() (int, error) {
	return u.repo.GetTotalCount()
}
