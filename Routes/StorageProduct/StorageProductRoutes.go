package StorageProduct

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
)

func ConfigureStorageProductRoutes(router *gin.RouterGroup, controller *Controllers.StorageProductController) {
	routerStorProd := router.Group("storage_product")

	//Methods GET
	routerStorProd.GET("/get/products/:id", controller.GetProductsByStorageId)
	routerStorProd.POST("/add/product/", controller.AddProductToStorage)
	routerStorProd.PATCH("/extract/product/", controller.ExtractProduct)

}
