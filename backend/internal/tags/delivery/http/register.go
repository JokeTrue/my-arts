package http

import (
	"github.com/JokeTrue/my-arts/internal/tags"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, useCase tags.UseCase) {
	handler := NewHandler(useCase)

	group := router.Group("/tags")
	{
		group.GET("/list", handler.GetTags)
		group.GET("/list/:product_id", handler.GetProductTags)
	}
}
