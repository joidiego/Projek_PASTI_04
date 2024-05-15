package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/joidiego/PASTI_KEl04/Config"
	controller "github.com/joidiego/PASTI_KEl04/Controllers"
	"github.com/joidiego/PASTI_KEl04/repository"
	"github.com/joidiego/PASTI_KEl04/service"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = config.SetupDatabaseConnection()
	sliderRepository repository.SliderRepository = repository.NewSliderRepository(db)
	sliderService    service.SliderService       = service.NewSliderService(sliderRepository)
	sliderController controller.SliderController = controller.NewSliderController(sliderService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	sliderRoutes := r.Group("/api/slider")
	{
		sliderRoutes.GET("/", sliderController.All)
		sliderRoutes.POST("/", sliderController.Insert)
		sliderRoutes.GET("/:id", sliderController.FindByID)
		sliderRoutes.PUT("/:id", sliderController.Update)
		sliderRoutes.DELETE("/:id", sliderController.Delete)
	}
	r.Run(":8009")
}
