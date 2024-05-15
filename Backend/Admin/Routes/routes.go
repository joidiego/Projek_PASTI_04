package routes

import (
	controllers "Admin/Controllers"
	middlewares "Admin/Middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	endpoint := app.Group("/admin")
	endpoint.Post("/login", controllers.LoginAdmin)
	endpoint.Use(middlewares.Middleware())
	endpoint.Get("/profile", controllers.GetProfile)
	endpoint.Post("/logout", controllers.LogouAdmin)
}
