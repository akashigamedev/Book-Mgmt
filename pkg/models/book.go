package models

import (
	"github.com/akashigamedev/book-mgmt/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() error {
	result := db.Create(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Book) UpdateBook() error {
	result := db.Save(b)
	return result.Error
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id uint64) (*Book, error) {
	var book Book
	result := db.Where("ID=?", id).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func DeleteBook(id uint64) error {
	return db.Delete(&Book{}, id).Error
}
