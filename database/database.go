package database

import (
	"belajar-golang-dasar/pkg/env"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var databaseInstance *gorm.DB

func InitializeDB() {
	fmt.Println("Loading database connection...")
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	databaseInstance = db
	fmt.Printf("Success connected to: %s\n", db.Name())
}

func GetDBInstance() *gorm.DB {
	return databaseInstance
}

func connectDB() (*gorm.DB, error) {
	// without parseTime=true, MySQL sends DATETIME values as raw []uint8 instead of converting them to time.Time
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.GetEnv("DB_USER"),
		env.GetEnv("DB_PASS"),
		env.GetEnv("DB_HOST"),
		env.GetEnv("DB_PORT"),
		env.GetEnv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
