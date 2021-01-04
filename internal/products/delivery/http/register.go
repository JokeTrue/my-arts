package http

import (
	"github.com/JokeTrue/my-arts/internal/products"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, useCase products.UseCase) {
	handler := NewHandler(useCase)

	group := router.Group("/products")
	{
		group.GET("/list", handler.GetProducts)
		group.POST("/product", handler.NotImplemented)
		group.GET("/product/:product_id", handler.GetProduct)
		group.PUT("/product/:product_id", handler.NotImplemented)
		group.DELETE("/product/:product_id", handler.DeleteProduct)
	}
}
