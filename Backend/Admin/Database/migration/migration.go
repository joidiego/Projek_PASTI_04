package migration

import (
	database "Admin/Database"
	"Admin/Models/entity"
	"fmt"
	"log"
)

func Migration() {

	err := database.DB.AutoMigrate(&entity.Admin{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
