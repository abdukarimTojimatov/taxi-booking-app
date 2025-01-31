package main

import (
	"context"
	"taxi-booking-app/config"
	"taxi-booking-app/routes"
	"taxi-booking-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Environment variables ni yuklash
	config.LoadEnv()

	// MongoDB ga ulanish
	client, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("taxi_booking")

	// Gin server yaratish
	r := gin.Default()

	// Servicelarni yaratish
	userService := services.NewUserService(db)
	driverService := services.NewDriverService(db)
	orderService := services.NewOrderService(db)

	// Routelarni ulash
	routes.UserRoutes(r, userService)
	routes.DriverRoutes(r, driverService)
	routes.OrderRoutes(r, orderService)

	// Serverni ishga tushirish
	r.Run(":" + config.GetPort())
}