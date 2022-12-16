package models

import (
	"github.com/ibnuyezid/go-bookstore/pkg/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
func GetAllBooks() []Book {
	var book []Book
	db.Find(&book)
	return book
}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("ID=?", Id).Find(&getbook)
	return &getbook, db
}
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
