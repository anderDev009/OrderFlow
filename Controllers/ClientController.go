package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"orderflow.com/v2/contracts/services"
	"orderflow.com/v2/dto/client"
)

type ClientController struct {
	clientService services.IClientService
}

func NewClientController(clientService services.IClientService) *ClientController {
	return &ClientController{clientService: clientService}
}

func (c *ClientController) GetClientList(ctx *gin.Context) {
	clients, err := c.clientService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, clients)
}
func (c *ClientController) GetClientById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	client, err := c.clientService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if client.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
	}
	ctx.JSON(http.StatusOK, client)
}
func (c *ClientController) CreateClient(ctx *gin.Context) {
	var add client.DtoAdd
	if err := ctx.ShouldBind(&add); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := c.clientService.Create(&add)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, add)
}
func (c *ClientController) UpdateClient(ctx *gin.Context) {
	var update client.DtoUpdate
	if err := ctx.ShouldBind(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := c.clientService.Update(&update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, update)
}
func (c *ClientController) DeleteClient(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = c.clientService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusNoContent, nil)
}
