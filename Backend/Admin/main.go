package main

import (
	"Admin/Database/migration"
	database "Admin/Database/seeders"
	routes "Admin/Routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	migration.Migration()
	// seeders.SeederData()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	routes.SetUp(app)

	err := app.Listen(":8001")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

}
