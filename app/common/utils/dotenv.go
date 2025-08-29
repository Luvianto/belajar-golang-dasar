package commonutils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Loading environment variables...")
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file:" + err.Error())
	}
	fmt.Println("Success loading .env file")
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable" + key + " not set")
	}
	return value
}
