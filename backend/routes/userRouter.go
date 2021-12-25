package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(icomingRoutes *gin.Engine) {
	icomingRoutes.get("/users", controller.GetUsers())
	icomingRoutes.get("/users/:user_id", controller.GetUser())
	icomingRoutes.POST("/users/signup", controller.SignUp())
	icomingRoutes.POST("/users/login", controller.Login())

}
