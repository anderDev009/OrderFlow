package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Dto/StorageProductDtos"
	"orderflow.com/v2/contracts/services"
)

type StorageProductController struct {
	service services.IStorageProductService
}

func NewStorageProductController(service services.IStorageProductService) *StorageProductController {
	return &StorageProductController{
		service: service,
	}
}

func (controller *StorageProductController) AddProductToStorage(ctx *gin.Context) {
	//validando el json que envia
	var dtoAdd StorageProductDtos.DtoAdd
	if err := ctx.ShouldBind(&dtoAdd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//llamando al sevicio
	err := controller.service.Create(&dtoAdd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": dtoAdd})
}
func (controller *StorageProductController) ExtractProduct(ctx *gin.Context) {
	var extractProd StorageProductDtos.ExtractProductDto
	if err := ctx.ShouldBind(&extractProd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//llamando al servicio
	errService := controller.service.ExtractProduct(extractProd.StorageId, extractProd.Quantity)
	if errService != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errService.Error()})
	}
	ctx.JSON(http.StatusNoContent, nil)
	return
}
func (controller *StorageProductController) GetProductsByStorageId(ctx *gin.Context) {
	var storageId = ctx.Param("id")
	idParsed, errParse := strconv.ParseUint(storageId, 10, 64)
	if errParse != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errParse.Error()})
		return
	}
	products, err := controller.service.GetProductsByStorageId(uint(idParsed))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
