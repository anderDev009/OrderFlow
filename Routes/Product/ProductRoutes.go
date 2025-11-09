package Product

import (
	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Controllers"
)

func ConfigureProductRoutes(server *gin.RouterGroup, controller *Controllers.ProductController) {
	productGroup := server.Group("product")
	//methods GET
	productGroup.GET("", controller.ListProducts)
	productGroup.GET("client/:id", controller.GetProductListWithClientId)
	productGroup.GET("/get/:id", controller.GetProductById)
	//CREATE and UPDATE
	productGroup.POST("", controller.AddProduct)
	productGroup.PUT("", controller.UpdateProduct)
	//DELETE
	productGroup.DELETE("/:id", controller.DeleteProduct)
}
