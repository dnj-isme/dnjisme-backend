package middleware

import (
	"net/http"

	"dnj-backend/config"

	"github.com/gin-gonic/gin"
)

func TestOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.LoadEnv().Mode != "test" {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "forbidden",
				"message": "This route is only accessible in test mode",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
