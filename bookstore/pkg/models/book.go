package models

import (
	"github.com/iamvibs/go-projects/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) UpdateBook() *Book {
	db.Save(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).First(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.First(&book, Id) // Fetch the book by ID
	db.Delete(&book)    // Delete the fetched book
	return book
}
