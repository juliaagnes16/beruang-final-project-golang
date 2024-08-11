package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(postgres.Open(databaseInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{}, &models.Expense{})

	DB = database
}
