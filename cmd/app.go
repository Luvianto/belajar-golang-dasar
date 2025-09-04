package cmd

import (
	"belajar-golang-dasar/app/api/models"
	"belajar-golang-dasar/app/api/services"
	"belajar-golang-dasar/app/config/database"
	"belajar-golang-dasar/app/config/repository/mysql"
	"fmt"
)

func App() {
	database.InitializeDB()
	db := database.GetDBInstance()

	// userRepo := mysql.NewUserRepository(db)
	// userService := services.NewUserService(userRepo)

	memberRep := mysql.NewMemberRepository(db)
	memberService := services.NewMemberService(memberRep)

	userCreate := &models.UserCreate{
		IsAdmin:  false,
		Email:    "John@gmail.com",
		Password: "password123",
		Phone:    "08123456789",
	}

	memberCreate := &models.MemberCreate{
		User:              *userCreate,
		Name:              "John Doe",
		Major:             "Computer Science",
		ProfilePictureUrl: "",
	}

	member, err := memberService.CreateMember(memberCreate)

	fmt.Println(member)
	fmt.Println(err)
}
