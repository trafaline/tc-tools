package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientKey := c.GetHeader("X-API-KEY")

		serverKey := os.Getenv("API_KEY")

		if clientKey == "" || clientKey != serverKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Invalid or missing API Key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
