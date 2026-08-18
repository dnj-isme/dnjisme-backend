package controllers

import (
	"dnj-backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	status := "running"
	dbStatus := "connected"

	sqlDB, err := database.DB.DB()
	if err != nil {
		status = "error"
		dbStatus = "unavailable"
	} else if err := sqlDB.Ping(); err != nil {
		status = "error"
		dbStatus = "disconnected"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "API is healthy",
		"status":   status,
		"database": dbStatus,
	})
}
