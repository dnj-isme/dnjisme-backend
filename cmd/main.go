package main

import (
	"dnj-backend/config"
	"dnj-backend/database"
	"dnj-backend/routes"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	migrateFlag := flag.Bool("migrate", false, "Run database migrations and seed data")
	flag.Parse()

	cfg := config.LoadEnv()
	database.ConnectDatabase(cfg.DBDSN)

	if *migrateFlag {
		// Migrate Data
		database.Migrate(database.DB)
		log.Println("Migration Done!")
	} else {
		// Run Server
		port := cfg.Port
		if port == "" {
			port = "8080"
		}

		r := gin.Default()
		routes.RegisterRoutes(r)
		r.Run(":" + port)
	}
}
