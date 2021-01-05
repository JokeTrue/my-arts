package http

import (
	"github.com/JokeTrue/my-arts/internal/categories"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, useCase categories.UseCase) {
	handler := NewHandler(useCase)

	group := router.Group("/categories")
	{
		group.GET("/list", handler.GetCategories)
	}
}
