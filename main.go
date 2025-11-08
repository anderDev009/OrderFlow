package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	//corriendo el servidor
	err := server.Run()
	if err != nil {
		panic("Failed to start the server")
	}
}
