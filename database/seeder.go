package database

import (
	"belajar-golang-dasar/app/api/models"
	commonutils "belajar-golang-dasar/common/utils"
	"fmt"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	InitAdmin(db)
	InitMember(db)
}

func InitAdmin(db *gorm.DB) {
	email := commonutils.GetEnv("USER_ADMIN_EMAIL")

	var isExists bool
	query := db.Table("users").Select("count(*) > 0").Where("email = ?", email).Find(&isExists)
	if query.Error != nil {
		fmt.Println("Error:", query.Error)
		return
	}

	if !isExists {
		fmt.Println("Loading seeder admin...")
		isAdmin := true
		uuid := commonutils.GenerateUUID()
		email := commonutils.GetEnv("USER_ADMIN_EMAIL")
		password := commonutils.GetEnv("USER_ADMIN_PASSWORD")
		phone := commonutils.GetEnv("USER_ADMIN_PHONE")
		encryptedPassword, err := commonutils.Encrypt(&password)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		db.Exec(`
		INSERT INTO users
			(uuid, is_admin, email, password, phone, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, now(), now())`, uuid, isAdmin, email, encryptedPassword, phone)
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
		var user models.User
		selectResult := db.Table("users").
			Select("UUID").
			Where("email = ?", commonutils.GetEnv("USER_ADMIN_EMAIL")).
			First(&user)
		if selectResult.Error != nil {
			fmt.Println("Error:", selectResult.Error)
			return
		}

		name := "Luvianto"
		major := "Computer Science"
		profilePictureUrl := ""

		db.Exec(`
		INSERT INTO members
			(user_id, name, major, profile_picture_url, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, now(), now())`,
			user.UUID, name, major, profilePictureUrl)

		fmt.Println("Success seeding member Luvianto")
	}
}
