package routes

import (
	"dnj-backend/controllers/swgoh"
	"dnj-backend/middleware"

	"github.com/gin-gonic/gin"
)

func registerSWGoHRoutes(router *gin.Engine) {
	route := router.Group("swgoh")

	grandivory := route.Group("grandivory")
	{
		grandivory.GET("fetch-player/:allycode", swgoh.FetchPlayerData)
		grandivory.GET("fetch-characters", swgoh.FetchAllCharacters)
		grandivory.GET("fetch-character", swgoh.FetchAllCharacters)
	}
	modTools := route.Group("mod-template")
	{
		modTools.POST("", middleware.MustLogin(), swgoh.SaveTemplate)
		modTools.DELETE("", middleware.MustLogin(), swgoh.DeleteTemplate)
		modTools.GET("", middleware.MustLogin(), swgoh.FetchTemplates)
	}
}
