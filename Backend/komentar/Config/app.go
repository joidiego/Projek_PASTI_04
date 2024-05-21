package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/service_comment?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = conn
	fmt.Println("Connected to database")

	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is nil. Did you forget to call Connect()?")
	}
	return db
}
