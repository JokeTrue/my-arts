package http

import (
	"github.com/JokeTrue/my-arts/internal/reviews"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, useCase reviews.UseCase) {
	handler := NewHandler(useCase)

	group := router.Group("/reviews")
	{
		group.POST("", handler.Create)
		group.PUT("/review/:review_id", handler.Update)
		group.GET("/review/:review_id", handler.GetReview)
		group.DELETE("/review/:review_id", handler.Delete)
		group.GET("/list/:user_id", handler.GetUserReviews)
	}
}
