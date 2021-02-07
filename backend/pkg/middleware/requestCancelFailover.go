package middleware

import "github.com/gin-gonic/gin"

func RequestCancelRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Request.Context().Done()
			}
		}()
		c.Next()
	}
}
