package routes

import (
	controllers "Product/Controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	product := app.Group("/product")
	product.Get("/", controllers.GetAllProduct)
	product.Post("/create", controllers.CreateProduct)
	product.Get("/:id", controllers.GetProductById)
	product.Put("/:id/edit", controllers.UpdateProduct)
	product.Delete("/:id/delete", controllers.DeleteProduct)
	product.Static("/image", controllers.PathImageProduct)
}
