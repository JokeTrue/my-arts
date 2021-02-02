package mysql

import (
	"database/sql"
	"github.com/JokeTrue/my-arts/internal/friendship"
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/tags"
	"github.com/jmoiron/sqlx"
	"sort"
	"strings"
)

type FriendshipRepository struct {
	db *sqlx.DB
}

func NewFriendshipRepository(db *sqlx.DB) *FriendshipRepository {
	return &FriendshipRepository{db: db}
}

func (r *FriendshipRepository) Get(id int) (*models.FriendshipRequest, error) {
	var request models.FriendshipRequest
	if err := r.db.Get(&request, QueryGetReviewByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, friendship.ErrFriendshipNotFound
		}
		return nil, friendship.ErrFriendshipQuery
	}
	return &request, nil
}

func (r *FriendshipRepository) Accept(user1, user2 int) error {
	ids := []int{user1, user2}
	sort.Ints(ids)

	result, err := r.db.Exec(QueryAcceptFriendshipRequest, ids[0], ids[1])
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return friendship.ErrFriendshipAlreadyExists
		}
		return friendship.ErrFriendshipQuery
	}

	_, err = result.LastInsertId()
	if err != nil {
		return friendship.ErrFriendshipQuery
	}

	return nil
}

func (r *FriendshipRepository) Delete(id int) error {
	res, err := r.db.Exec(QueryDeleteFriendshipRequest, id)
	if err != nil {
		return tags.ErrTagQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return friendship.ErrFriendshipQuery
	}
	if affect == 0 {
		return friendship.ErrFriendshipNotFound
	}

	return nil
}

func (r *FriendshipRepository) Create(request models.FriendshipRequest) (int, error) {
	result, err := r.db.Exec(QueryCreateFriendshipRequest, request.ActorID, request.FriendID)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, friendship.ErrFriendshipRequestAlreadyExists
		}
		return 0, friendship.ErrFriendshipQuery
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, friendship.ErrFriendshipQuery
	}

	return int(lastID), nil
}

func (r *FriendshipRepository) GetList(userId int) ([]*models.FriendshipRequest, error) {
	list := []*models.FriendshipRequest{}
	if err := r.db.Select(&list, QueryGetUserFriendshipRequests, userId); err != nil {
		return nil, friendship.ErrFriendshipQuery
	}
	return list, nil
}
