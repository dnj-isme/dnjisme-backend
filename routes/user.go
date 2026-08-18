package routes

import (
	"dnj-backend/controllers/user"
	"dnj-backend/middleware"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.Engine) {
	auth := router.Group("auth")
	{
		auth.POST("register", middleware.TestOnly(), user.RegisterUser)
		auth.POST("login", user.LoginUser)
		auth.POST("test-auth", middleware.TestOnly(), user.TestBearer)
	}
}
