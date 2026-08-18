package database

import (
	"dnj-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	for _, roleName := range []string{"admin", "user"} {
		var role models.Role
		err := db.Where("role_name = ?", roleName).First(&role).Error
		if err == nil {
			continue
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}
		if err := db.Create(&models.Role{RoleName: roleName}).Error; err != nil {
			return err
		}
	}
	return nil
}
