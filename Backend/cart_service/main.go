package main

import (
	config "cart_service/Config"
	controller "cart_service/Controllers"
	service "cart_service/Service"
	"cart_service/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	cartRepository repository.CartRepository = repository.NewCartRepository(db)
	CartService    service.CartService       = service.NewCartService(cartRepository)
	cartController controller.CartController = controller.NewCartController(CartService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	cartRoutes := r.Group("/api/carts")
	{
		cartRoutes.GET("/", cartController.All)
		cartRoutes.POST("/create", cartController.Insert)
		cartRoutes.PUT("/:id", cartController.Update)
		cartRoutes.DELETE("/:id", cartController.Delete)
	}
	r.Run(":8008")
}