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
	orders.GET("/get/detail/:id", orderController.GetOrderWithDetails)
	//methods POST
	orders.POST("", orderController.AddOrder)
	orders.POST("/add/detail", orderController.AddOrderDetail)
	//UPDATE
	orders.PUT("", orderController.UpdateOrder)
	//DELETE
	orders.DELETE("/:id", orderController.DeleteOrder)
	orders.DELETE("/detail/:id", orderController.DeleteOrder)
}
