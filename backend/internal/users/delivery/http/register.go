package http

import (
	"github.com/JokeTrue/my-arts/internal/users"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(authRouter *gin.RouterGroup, router *gin.Engine, useCase users.UseCase) {
	handler := NewHandler(useCase)

	group := authRouter.Group("/users")
	{
		group.GET("/me", handler.GetCurrentUser)
		group.GET("/search", handler.SearchUsers)
		group.GET("/user/:user_id", handler.GetUser)
		group.GET("/user/:user_id/friends", handler.GetUserFriends)
		group.GET("/total_count", handler.GetUsersTotalCount)
	}

	router.POST("/api/sign_up", handler.CreateUser)
}
