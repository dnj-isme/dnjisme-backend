package user

import (
	"dnj-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestBearer(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	claims, err := utils.ValidateBearerToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}

	c.Set("user_id", claims.UserID)
	c.Set("email", claims.Email)
	c.Set("role", claims.Role)

	c.JSON(http.StatusOK, gin.H{
		"message": "Bearer token is valid",
		"user_id": claims.UserID,
		"email":   claims.Email,
		"role":    claims.Role,
	})
}
