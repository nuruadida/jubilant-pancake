package db

import (
	"Nurul-trial/models"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		models.GetItGuys{},
		models.User{},
	)

	if err != nil {
		return
	}
}
