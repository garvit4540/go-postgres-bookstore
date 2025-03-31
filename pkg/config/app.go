package config

import (
	"github.com/garvit4540/go-postgres-bookstore/pkg/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=bookstore_user password=12345678 dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		trace.PushTrace(trace.ErrorFailedToConnectToPostgres, map[string]interface{}{"error": err})
		panic(err)
	}
	trace.PushTrace(trace.SuccessfullyConnectedToDatabase, nil)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
