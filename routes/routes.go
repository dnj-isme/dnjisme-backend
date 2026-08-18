package routes

import (
	"dnj-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)
	registerUserRoutes(router)
	registerSWGoHRoutes(router)
	registerPortfolioRoutes(router)
}
