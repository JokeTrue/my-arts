package usecase

import (
	"github.com/JokeTrue/my-arts/internal/friendship"
	"github.com/JokeTrue/my-arts/internal/models"
)

type FriendshipUseCase struct {
	repo friendship.Repository
}

func NewFriendshipUseCase(repo friendship.Repository) *FriendshipUseCase {
	return &FriendshipUseCase{repo: repo}
}

func (u *FriendshipUseCase) Get(id int) (*models.FriendshipRequest, error) {
	return u.repo.Get(id)
}

func (u *FriendshipUseCase) Accept(user1, user2 int) error {
	return u.repo.Accept(user1, user2)
}

func (u *FriendshipUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *FriendshipUseCase) Create(request models.FriendshipRequest) (int, error) {
	return u.repo.Create(request)
}

func (u *FriendshipUseCase) GetList(userId int) ([]*models.FriendshipRequest, error) {
	return u.repo.GetList(userId)
}
