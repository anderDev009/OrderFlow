package main

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
	"orderflow.com/v2/Repositories"
	"orderflow.com/v2/Routes"
	"orderflow.com/v2/Services"
	"orderflow.com/v2/database"
)

func main() {
	server := gin.Default()
	//database object
	ctx := database.GetContext()
	//dependency injection
	//repositories
	clientRepository := Repositories.NewClientRepository(ctx)
	productRepository := Repositories.NewProductRepository(ctx)
	customerRepository := Repositories.NewCustomerRepository(ctx)
	storageRepository := Repositories.NewStorageRepository(ctx)
	storageProductRepository := Repositories.NewStorageProductRepository(ctx)
	orderRepository := Repositories.NewOrderRepository(ctx)
	//services
	clientService := Services.NewClientService(clientRepository)
	productService := Services.NewProductService(productRepository)
	customerService := Services.NewCustomerService(customerRepository)
	storageService := Services.NewStorageService(storageRepository)
	storageProductService := Services.NewStorageProductService(storageProductRepository)
	orderService := Services.NewOrderService(orderRepository)
	//configurando controladores
	clientController := Controllers.NewClientController(clientService)
	productController := Controllers.NewProductController(productService)
	customerController := Controllers.NewCustomerController(customerService)
	storageControler := Controllers.NewStorageController(storageService)
	storageProductController := Controllers.NewStorageProductController(storageProductService)
	orderController := Controllers.NewOrderController(orderService)
	//route v1 config
	Routes.ConfigureV1(server,
		clientController,
		productController,
		customerController,
		storageControler,
		storageProductController,
		orderController)
	//running server
	err := server.Run()
	if err != nil {
		panic("Failed to start the server")
	}
}
