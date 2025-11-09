package Routes

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
	"orderflow.com/v2/Routes/Client"
	"orderflow.com/v2/Routes/Customer"
	"orderflow.com/v2/Routes/Product"
)

func ConfigureV1(server *gin.Engine,
	clientController *Controllers.ClientController,
	productController *Controllers.ProductController,
	customerController *Controllers.CustomerController) {
	routeV1 := server.Group("api/v1")
	//client routes
	Client.ConfigureClientRoutesV1(routeV1, clientController)
	Product.ConfigureProductRoutes(routeV1, productController)
	Customer.ConfigureCustomerRoutes(routeV1, customerController)
}
