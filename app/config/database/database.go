package database

import (
	commonutils "belajar-golang-dasar/app/common/utils"
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
		commonutils.GetEnv("DB_USER"),
		commonutils.GetEnv("DB_PASS"),
		commonutils.GetEnv("DB_HOST"),
		commonutils.GetEnv("DB_PORT"),
		commonutils.GetEnv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
