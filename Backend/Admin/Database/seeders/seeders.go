package seeders

import (
	database "Admin/Database"
	"Admin/Models/entity"
	utils "Admin/Utils"
	"fmt"
	"log"
)

func SeederData() {
	password, err := utils.GeneratePassword("destinasalto")
	if err != nil {
		log.Fatalf(err.Error())
	}

	admin := &entity.Admin{
		Username: "destinalapar",
		Password: password,
	}

	if err := database.DB.Create(&admin); err != nil {
		log.Fatalf("Failed to create admin: %v", err)
	}

	fmt.Println("Data Seeded Successfully")
}
