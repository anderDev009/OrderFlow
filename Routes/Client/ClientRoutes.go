package Client

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
)

func ConfigureClientRoutesV1(server *gin.RouterGroup, controller *Controllers.ClientController) {
	clientGroup := server.Group("/client")
	//Methods GET
	clientGroup.GET("/", controller.GetClientList)
	clientGroup.GET("/get/:id", controller.GetClientById)
	//Create And Delete
	clientGroup.POST("/", controller.CreateClient)
	clientGroup.PUT("/", controller.UpdateClient)
	//Delete
	clientGroup.DELETE("/:id", controller.DeleteClient)
}
