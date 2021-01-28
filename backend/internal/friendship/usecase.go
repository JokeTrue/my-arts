package friendship

import "github.com/JokeTrue/my-arts/internal/models"

type UseCase interface {
	Delete(id int) error
	Accept(user1, user2 int) error
	Get(id int) (*models.FriendshipRequest, error)
	Create(request models.FriendshipRequest) (int, error)
	GetList(userId int) ([]*models.FriendshipRequest, error)
}
