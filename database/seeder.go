package database

import (
	commonUtils "belajar-golang-dasar/common/utils"
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	userEntity "belajar-golang-dasar/internal/module/user/entity"
	"belajar-golang-dasar/pkg/env"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	InitAdmin(db)
	InitMember(db)
}

func InitAdmin(db *gorm.DB) {
	email := env.GetEnv("USER_ADMIN_EMAIL")

	var isExists bool
	query := db.Table("users").Select("count(*) > 0").Where("email = ?", email).Find(&isExists)
	if query.Error != nil {
		fmt.Println("Error:", query.Error)
		return
	}

	if !isExists {
		fmt.Println("Loading seeder admin...")
		isAdmin := true
		uuid := commonUtils.GenerateUUID()
		password := env.GetEnv("USER_ADMIN_PASSWORD")
		phone := env.GetEnv("USER_ADMIN_PHONE")
		encryptedPassword, err := commonUtils.Encrypt(&password)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		user := userEntity.User{
			UUID:     uuid,
			IsAdmin:  isAdmin,
			Email:    email,
			Password: encryptedPassword,
			Phone:    phone,
		}

		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed user: %v", err)
		}
		fmt.Println("Success seeding user admin")
	}
}

func InitMember(db *gorm.DB) {
	fmt.Println("Loading seeder member...")

	var isExists bool

	result := db.Table("members").
		Select("count(*) > 0").
		Where("name = ?", "Luvianto").
		Scan(&isExists)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return
	}

	if !isExists {
		var user userEntity.User
		selectResult := db.Table("users").
			Select("UUID").
			Where("email = ?", env.GetEnv("USER_ADMIN_EMAIL")).
			First(&user)
		if selectResult.Error != nil {
			fmt.Println("Error:", selectResult.Error)
			return
		}

		member := memberEntity.Member{
			UserID:            user.UUID,
			Name:              "Luvianto",
			Major:             "Computer Science",
			ProfilePictureUrl: "",
		}

		if err := db.Create(&member).Error; err != nil {
			log.Fatalf("Failed to seed member: %v", err)
		}

		fmt.Println("Success seeding member Luvianto")
	}
}
