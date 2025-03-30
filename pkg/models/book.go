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
		trace.PushTrace(trace.ErrorMigratingDatabases, nil)
		panic(err)
	}
}
