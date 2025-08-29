package main

import (
	"belajar-golang-dasar/app/api/models"
	"belajar-golang-dasar/app/api/services"
	commonutils "belajar-golang-dasar/app/common/utils"
	"belajar-golang-dasar/app/config/database"
	"belajar-golang-dasar/app/config/repository/mysql"
	"flag"
	"fmt"
)

func main() {
	commonutils.LoadEnv()

	runSeeder := flag.Bool("seed", false, "Menjalanakan seeder")
	runMigration := flag.Bool("migrate", false, "Menjalankan migration")
	flag.Parse()

	if *runMigration {
		migration()
	}

	if *runSeeder {
		seeder()
	}

	if !*runMigration && !*runSeeder {
		app()
	}
}

func app() {
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

func migration() {
	database.InitializeDB()
	db := database.GetDBInstance()

	fmt.Println("Loading migrating database...")
	err := database.Migrate(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success migrating database")
}

func seeder() {
	database.InitializeDB()
	db := database.GetDBInstance()
	fmt.Println("Seeding database...")
	database.Seeder(db)
}
