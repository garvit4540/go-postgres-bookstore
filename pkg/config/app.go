package config

import (
	"github.com/garvit4540/go-mysql-bookstore/pkg/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "bookstore_user:12345678@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		trace.PushTrace(trace.ErrorFailedToConnectToMySql, map[string]interface{}{"error": err})
		panic(err)
	}
	trace.PushTrace(trace.SuccessfullyConnectedToDatabase, nil)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
