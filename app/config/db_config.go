package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetDbConfig() map[string]interface{} {
	// loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	return map[string]interface{}{
		"DB_HOST":     host,
		"DB_PORT":     port,
		"DB_USER":     user,
		"DB_PASSWORD": password,
		"DB_NAME":     dbname,
	}
}
