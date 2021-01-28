package http

import (
	"github.com/JokeTrue/my-arts/internal/friendship"
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckObjectPermissions(c *gin.Context, useCase friendship.UseCase) (*models.FriendshipRequest, error) {
	rawRequestId := c.Param("request_id")
	requestId, err := strconv.Atoi(rawRequestId)
	if err != nil {
		return nil, err
	}

	request, err := useCase.Get(requestId)
	if err != nil {
		return nil, err
	}

	userId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		return nil, err
	}

	if request.FriendID != userId {
		return nil, friendship.ErrFriendshipPermissionDenied
	}

	return request, nil
}
