package http

import (
	"github.com/JokeTrue/my-arts/internal/friendship"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, useCase friendship.UseCase) {
	handler := NewHandler(useCase)

	group := router.Group("/friendship")
	{
		group.GET("/list/:user_id", handler.GetUserFriendshipRequests)
		group.POST("/create/:user_id", handler.CreateFriendshipRequest)
		group.POST("/action/:action/:request_id", handler.ActionFriendshipRequest)
	}
}
