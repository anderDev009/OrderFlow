package Order

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
)

func ConfigureRoutes(router *gin.RouterGroup, orderController *Controllers.OrderController) {
	orders := router.Group("/orders")
	//methos GET
	orders.GET("/get/client/:id", orderController.GetOrdersByClientId)
	orders.GET("/get/:id", orderController.GetOrdersByClientId)
	//methods POST
	orders.POST("", orderController.AddOrder)
	//UPDATE
	orders.PUT("", orderController.UpdateOrder)
	//DELETE
	orders.DELETE("/:id", orderController.DeleteOrder)
}
