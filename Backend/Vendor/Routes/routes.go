package routes

import (
	controllers "Vendor/Controllers"
	middleware "Vendor/Middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	Vendor := app.Group("/Vendor")
	Vendor.Post("/register", controllers.VendorRegister)
	Vendor.Post("/login", controllers.VendorLogin)
	Vendor.Use(middleware.Middleware())
	Vendor.Get("/profile", controllers.VendorGetProfile)
	Vendor.Post("/logout", controllers.VendorLogout)
}
