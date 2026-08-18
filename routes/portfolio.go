package routes

import (
	"dnj-backend/controllers/portfolio"

	"github.com/gin-gonic/gin"
)

func registerPortfolioRoutes(router *gin.Engine) {
	router.GET("/portfolio", portfolio.GetPortfolio)
}
