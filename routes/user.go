package routes

import (
	"taxi-booking-app/controllers"
	"taxi-booking-app/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userService services.UserService) {
	userController := controllers.NewUserController(userService)

	user := r.Group("/user")
	{
		user.POST("/register", userController.RegisterUser)
		user.POST("/login", userController.LoginUser)
	}
}