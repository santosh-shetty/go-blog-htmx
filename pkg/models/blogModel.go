package models

import (
	"github.com/santosh-shetty/blog/pkg/config"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title       string `json:"title"`
	ShortDesc   string `json:"short_desc"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Category    int64  `json:"category"`
}

func FindAll() ([]Blog, error) {
	var blogs []Blog
	result := config.DB.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}
func FindById(id int64) (Blog, error) {
	var blog Blog
	result := config.DB.First(&blog, id)
	if result.Error != nil {
		return Blog{}, result.Error
	}
	return blog, nil
}
func UpdateBlogById(id int64, blog Blog) error {
	result := config.DB.Model(&Blog{}).Where("id=?", id).Updates(blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func AddBlog(data Blog) error {
	result := config.DB.Create(&data)
	return result.Error
}

func DeleteBlog(id int64) error {
	var blog Blog
	result := config.DB.Unscoped().Delete(&blog, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SearchBlogbyTitle(title string) ([]Blog, error) {
	var blogs []Blog
	result := config.DB.Where("title LIKE ?", "%"+title+"%").Find(&blogs)
	if result.Error != nil {
		return []Blog{}, result.Error
	}
	return blogs, nil

}
