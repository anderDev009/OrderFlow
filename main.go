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
	//services
	clientService := Services.NewClientService(clientRepository)
	productService := Services.NewProductService(productRepository)
	//configurando controladores
	clientController := Controllers.NewClientController(clientService)
	productController := Controllers.NewProductController(productService)
	//route v1 config
	Routes.ConfigureV1(server, clientController, productController)
	//running server
	err := server.Run()
	if err != nil {
		panic("Failed to start the server")
	}
}
