package database

import (
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	userEntity "belajar-golang-dasar/internal/module/user/entity"

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
	&userEntity.User{},
	&memberEntity.Member{},
}
