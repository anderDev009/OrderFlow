package main

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
	"orderflow.com/v2/Repositories"
	"orderflow.com/v2/Routes"
	"orderflow.com/v2/Services"
)

func main() {
	server := gin.Default()
	//dependency injection
	//repositories
	clientRepository := Repositories.NewClientRepository()
	//services
	clientService := Services.NewClientService(clientRepository)
	//configurando controladores
	clientController := Controllers.NewClientController(clientService)
	//route v1 config
	Routes.ConfigureV1(server, clientController)
	//running server
	err := server.Run()
	if err != nil {
		panic("Failed to start the server")
	}
}
