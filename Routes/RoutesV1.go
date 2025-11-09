package Routes

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
	"orderflow.com/v2/Routes/Client"
)

func ConfigureV1(server *gin.Engine, clientController *Controllers.ClientController) {
	routeV1 := server.Group("api/v1")
	//client routes
	Client.ConfigureClientRoutesV1(routeV1, clientController)
}
