package routes

import (
	controllers "Customer/Controllers"
	middleware "Customer/Middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUp(App *fiber.App) {
	customer := App.Group("/customer")
	customer.Post("/register", controllers.RegistrationCustomer)
	customer.Post("/login", controllers.LoginCustomer)
	customer.Static("/image", controllers.PathImageCustomer)
	customer.Put("/forgot_password", controllers.ForgotPassword)
	customer.Use(middleware.CheckLogin())
	customer.Get("/profile", controllers.GetProfile)
	customer.Put("/updateProfile", controllers.UpdateProfile)
	customer.Put("/edit_password", controllers.EditPassword)
	customer.Post("/logout", controllers.CustomerLogout)

}
