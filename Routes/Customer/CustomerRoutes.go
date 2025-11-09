package Customer

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
)

func ConfigureCustomerRoutes(server *gin.RouterGroup, controller *Controllers.CustomerController) {
	customerGroup := server.Group("/customer")
	//methods GET
	customerGroup.GET("", controller.GetCustomers)
	customerGroup.GET("get/:id", controller.GetCustomer)
	customerGroup.GET("/client/:id", controller.GetCustomersWithClientId)
	customerGroup.POST("", controller.CreateCustomer)
	customerGroup.PUT("", controller.CreateCustomer)
	customerGroup.DELETE("/:id", controller.CreateCustomer)
}
