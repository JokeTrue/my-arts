package http

import (
	"github.com/JokeTrue/my-arts/internal/users"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, useCase users.UseCase) {
	handler := NewHandler(useCase)

	group := router.Group("/users")
	{
		group.GET("/me", handler.GetCurrentUser)
		group.POST("/sign_up", handler.CreateUser)
	}
}
