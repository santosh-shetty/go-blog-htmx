package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database:", err)
	}
	fmt.Println("DB Connected Sucessfully!")
	// migrateErr := database.AutoMigrate(new(models.Category))
	// if migrateErr != nil {
	// 	log.Println("Error During Auto Migrate DB", migrateErr)
	// }

	DB = database

}

func DBClose() {
	database, err1 := DB.DB()
	if err1 != nil {
		log.Println("Failed to get database connection :", err1)
	}
	err2 := database.Close()
	if err2 != nil {
		log.Println("Failed to close database connection:", err2)
	} else {
		fmt.Println("DB Connection Closed Successfully!")
	}
}
