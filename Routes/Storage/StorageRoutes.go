package Storage

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
)

func ConfigureStorageRoutes(router *gin.RouterGroup, storageController *Controllers.StorageController) {
	storageRouter := router.Group("storage")
	//methods GET
	storageRouter.GET("", storageController.GetAllStorages)
	storageRouter.GET("/client/:id", storageController.GetAllStoragesByClientId)
	storageRouter.GET("/get/:id", storageController.GetById)
	//methods POST
	storageRouter.POST("", storageController.CreateStorage)
	//methods PUT
	storageRouter.PUT("", storageController.UpdateStorage)
	storageRouter.DELETE("/:id", storageController.Delete)
}
