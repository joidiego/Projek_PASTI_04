package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DBo *gorm.DB

func Connect() {
	const con = "root@tcp(localhost)/service_admin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := con
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	const connect = "root@tcp(localhost)/service_order?charset=utf8&parseTime=True&loc=Local"
	ds := connect
	dbo, err := gorm.Open(mysql.Open(ds), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database")
	}

	DB = db

	DBo = dbo

	fmt.Println("Successfully Connect to database")
}
