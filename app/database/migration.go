package database

import (
	"Eatiplan-Auth/app/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
