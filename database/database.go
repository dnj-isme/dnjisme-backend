package database

import (
	"dnj-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB = database
}

func Migrate(db *gorm.DB) {
	// Auto migrate all models
	err := db.AutoMigrate(
		// General
		&models.Role{},
		&models.User{},
		&models.Portfolio{},

		// Grandivory
		&models.SquadUnit{},
		&models.Squad{},
		&models.SquadTemplate{},
	)
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	err2 := _prepareRole(db)
	if err2 != nil {
		panic("Failed to insert Role: " + err2.Error())
	}
}

func _prepareRole(db *gorm.DB) error {
	var roles = []models.Role{
		{RoleName: "admin"},
		{RoleName: "user"},
	}
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "role_name"}},
		DoNothing: true,
	}).Create(&roles).Error
}
