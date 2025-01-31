package routes

import (
	"taxi-booking-app/controllers"
	"taxi-booking-app/services"

	"github.com/gin-gonic/gin"
)

func DriverRoutes(r *gin.Engine, driverService services.DriverService) {
	driverController := controllers.NewDriverController(driverService)

	driver := r.Group("/driver")
	{
		driver.POST("/location", driverController.UpdateDriverLocation)
	}
}