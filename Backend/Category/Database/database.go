package database

import (
	models "Category/Models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	const con = "root@tcp(localhost)/service_category?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := con
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	DB = db

	fmt.Println("Successfully Connect to database")

	DB.AutoMigrate(&models.Category{})
}
