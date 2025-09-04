package database

import (
	"belajar-golang-dasar/app/api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	for _, model := range modelList {
		if err := db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}

	return nil
}

var modelList = []any{
	&models.User{},
	&models.Member{},
}
