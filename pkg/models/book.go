package models

import (
	"github.com/Jasperino64/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100)" json:"title"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	Publication string `gorm:"type:varchar(100)" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id uint) (*Book, *gorm.DB) {
	var book Book
	// db := db.Where("ID = ?", id).Find(&book)
	db := db.First(&book, id)
	return &book, db
}

// func (b *Book) UpdateBook(id int) *Book {
// 	db.Save(b)
// 	return b
// }

func DeleteBook(id uint) Book {
	var book Book
	db.Delete(&book, id)
	return book
}
