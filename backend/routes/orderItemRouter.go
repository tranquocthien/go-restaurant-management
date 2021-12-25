package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())

	incomingRoutes.POST("/orderitems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderitems/:orderitem_id", controller.UpdateOrderItem())
}
