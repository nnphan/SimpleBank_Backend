package middlewares

import (
	"simplebank/response"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.StatusInvalidToken, "Invalid token")
			c.Abort() // Stop further handlers from being executed
			return
		}
		c.Next() // Continue to the next handler
	}
}