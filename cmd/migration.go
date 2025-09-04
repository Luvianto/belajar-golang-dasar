package cmd

import (
	"belajar-golang-dasar/app/config/database"
	"fmt"
)

func Migration() {
	database.InitializeDB()
	db := database.GetDBInstance()

	fmt.Println("Loading migrating database...")
	err := database.Migrate(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success migrating database")
}
