package models

import (
	"github.com/garvit4540/go-mysql-bookstore/pkg/config"
	"github.com/garvit4540/go-mysql-bookstore/pkg/trace"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `gorm:"" json:"author"`
	Publication string `gorm:"" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		trace.PushTrace(trace.ErrorMigratingDatabases, map[string]interface{}{"error": err})
		panic(err)
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	newDbCon := db.Where("ID=?", id).Find(&getBook)
	return &getBook, newDbCon
}

func DeleteBook(id int64) *Book {
	book := &Book{}
	_ = db.Where("ID=?", id).Delete(book)
	return book
}
