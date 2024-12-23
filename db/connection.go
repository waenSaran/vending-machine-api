package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/waenSaran/vending-machine-api/app/config"
	"github.com/waenSaran/vending-machine-api/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	DB = OpenConnection()
}

func OpenConnection() *gorm.DB {
	config := config.GetDbConfig()
	host := config["DB_HOST"]
	port := config["DB_PORT"]
	user := config["DB_USER"]
	password := config["DB_PASSWORD"]
	dbname := config["DB_NAME"]

	// Add timezone to DSN
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=UTC",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Categories{}, &models.SubCategories{}, &models.Products{}, &models.Money{})

	return db
}
