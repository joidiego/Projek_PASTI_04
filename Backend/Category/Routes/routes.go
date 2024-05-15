package routes

import (
	controllers "Category/Controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUpCategory(app *fiber.App) {
	routes := app.Group("/category")
	routes.Get("/", controllers.GetAllCategory)
	routes.Post("/create", controllers.CreateCategory)
	routes.Get("/:id", controllers.GetCategoryById)
	routes.Put("/:id/edit", controllers.UpdateCategory)
	routes.Delete("/:id/delete", controllers.DeleteCategory)
}
