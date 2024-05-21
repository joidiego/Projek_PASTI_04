package routes

import (
	controllers "Order/Controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	orders := app.Group("/orders")
	orders.Post("/create", controllers.CreateOrder)
	orders.Get("/", controllers.GetAllOrder)
	orders.Get("/:id", controllers.GetOrderById)
}
