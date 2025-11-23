package Routes

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
	"orderflow.com/v2/Routes/Client"
	"orderflow.com/v2/Routes/Customer"
	"orderflow.com/v2/Routes/Order"
	"orderflow.com/v2/Routes/Product"
	"orderflow.com/v2/Routes/Storage"
	"orderflow.com/v2/Routes/StorageProduct"
)

func ConfigureV1(server *gin.Engine,
	clientController *Controllers.ClientController,
	productController *Controllers.ProductController,
	customerController *Controllers.CustomerController,
	storageController *Controllers.StorageController,
	storageProductController *Controllers.StorageProductController,
	orderController *Controllers.OrderController) {
	routeV1 := server.Group("api/v1")
	//client routes
	Client.ConfigureClientRoutesV1(routeV1, clientController)
	Product.ConfigureProductRoutes(routeV1, productController)
	Customer.ConfigureCustomerRoutes(routeV1, customerController)
	Storage.ConfigureStorageRoutes(routeV1, storageController)
	StorageProduct.ConfigureStorageProductRoutes(routeV1, storageProductController)
	//order routes
	Order.ConfigureRoutes(routeV1, orderController)
}
