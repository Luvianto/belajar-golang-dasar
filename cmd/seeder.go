package cmd

import (
	"belajar-golang-dasar/app/config/database"
	"fmt"
)

func Seeder() {
	database.InitializeDB()
	db := database.GetDBInstance()
	fmt.Println("Seeding database...")
	database.Seeder(db)
}
