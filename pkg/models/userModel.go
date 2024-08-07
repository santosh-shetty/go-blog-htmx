package models

import (
	"log"

	"github.com/santosh-shetty/blog/pkg/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user User) error {
	migrateErr := config.DB.AutoMigrate(new(User))
	if migrateErr != nil {
		log.Println("Error During Auto Migrate DB", migrateErr)
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUserByEmail(email string) (User, error) {
	var user User
	result := config.DB.Where("email=?", email).Find(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
