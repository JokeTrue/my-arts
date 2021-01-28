package http

import (
	"github.com/JokeTrue/my-arts/internal/friendship"
	"github.com/JokeTrue/my-arts/internal/models"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/JokeTrue/my-arts/pkg/jwt"
	"github.com/JokeTrue/my-arts/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase friendship.UseCase
}

func NewHandler(useCase friendship.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetUserFriendshipRequests(c *gin.Context) {
	rawUserId := c.Param("user_id")
	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		appErrors.JSONError(c, err, userId)
		return
	}

	requests, err := h.useCase.GetList(userId)
	if err != nil {
		appErrors.JSONError(c, err, userId)
		return
	}

	c.JSON(http.StatusOK, requests)
}

func (h *Handler) CreateFriendshipRequest(c *gin.Context) {
	rawFriendId := c.Param("user_id")
	friendId, err := strconv.Atoi(rawFriendId)
	if err != nil {
		appErrors.JSONError(c, err, rawFriendId)
		return
	}

	userId, err := jwt.GetCurrentUserID(c)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	request := models.FriendshipRequest{
		ActorID:  userId,
		FriendID: friendId,
	}
	request.ID, err = h.useCase.Create(request)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusCreated, request)
}

func (h *Handler) ActionFriendshipRequest(c *gin.Context) {
	request, err := CheckObjectPermissions(c, h.useCase)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	action := c.Param("action")
	allowedActions := []string{"accept", "decline"}
	if !utils.Contains(allowedActions, action) {
		err = friendship.ErrFriendshipUnknownAction
		appErrors.JSONError(c, err, nil)
		return
	}

	if action == "accept" {
		err = h.useCase.Accept(request.ActorID, request.FriendID)
		if err != nil {
			appErrors.JSONError(c, err, nil)
			return
		}
	}

	err = h.useCase.Delete(request.ID)
	if err != nil {
		appErrors.JSONError(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
