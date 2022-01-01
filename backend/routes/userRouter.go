package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(icomingRoutes *gin.Engine) {
	icomingRoutes.GET("/users", controller.GetUsers())
	icomingRoutes.GET("/users/:user_id", controller.GetUser())
	icomingRoutes.POST("/users/signup", controller.SignUp())
	icomingRoutes.POST("/users/login", controller.Login())

}
