package main

import (
	database "Product/Database"
	routes "Product/Routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New())

	routes.SetUp(app)

	err := app.Listen(":8005")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
