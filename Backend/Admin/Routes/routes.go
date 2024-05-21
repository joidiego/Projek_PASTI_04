package routes

import (
	controllers "Admin/Controllers"
	middleware "Admin/Middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	endpoint := app.Group("/admin")
	endpoint.Post("/login", controllers.LoginAdmin)
	endpoint.Use(middleware.Middleware())
	endpoint.Get("/profile", controllers.GetProfile)
	endpoint.Post("/logout", controllers.LogouAdmin)
	endpoint.Put("/approve/:id", controllers.DataOrder)
}
