package routes

import (
	"taxi-booking-app/controllers"
	"taxi-booking-app/services"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, orderService services.OrderService) {
	orderController := controllers.NewOrderController(orderService)

	order := r.Group("/order")
	{
		order.POST("/create", orderController.CreateOrder)
	}
}