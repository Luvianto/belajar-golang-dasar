package database

import (
	"belajar-golang-dasar/pkg/env"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databaseInstance *gorm.DB

func GetDBInstance() *gorm.DB {
	return databaseInstance
}

func InitializeDB() {
	fmt.Println("Loading database connection...")

	dbType := env.GetEnv("DB_TYPE")

	var err error

	switch dbType {
	case "mysql":
		databaseInstance, err = connectMySQL()
	case "sqlite":
		databaseInstance, err = connectSQLite()
	default:
		log.Fatalf("Unsupported DB_TYPE: %s", dbType)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("Success connected to: %s\n", databaseInstance.Name())
}

func connectSQLite() (*gorm.DB, error) {
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, nil
}

func connectMySQL() (*gorm.DB, error) {
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
