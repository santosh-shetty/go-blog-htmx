package models

import (
	"log"

	"github.com/santosh-shetty/blog/pkg/config"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AllCategory() ([]Category, error) {
	var categories []Category
	err := config.DB.Find(&categories)
	if err.Error != nil {
		return nil, err.Error
	}
	return categories, nil
}

func AddCategory(data Category) error {
	migrateErr := config.DB.AutoMigrate(new(Category))
	if migrateErr != nil {
		log.Println("Error During Auto Migrate DB", migrateErr)
	}
	err := config.DB.Create(&data)
	if err != nil {
		return err.Error
	}
	return nil
}

func DeleteCategory(id int64) error {
	var category Category
	err := config.DB.Unscoped().Delete(&category, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func CategoryById(id int64) (Category, error) {
	var category Category
	result := config.DB.First(&category, id)
	if result.Error != nil {
		return Category{}, result.Error
	}
	return category, nil
}

func UpdateCategoryById(id int64, updatedCategory Category) error {
	result := config.DB.Model(&Category{}).Where("id = ?", id).Updates(updatedCategory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
