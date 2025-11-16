package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Dto/StorageDtos"
	"orderflow.com/v2/contracts/services"
)

type StorageController struct {
	service services.IStorageService
}

func NewStorageController(service services.IStorageService) *StorageController {
	return &StorageController{service: service}
}

func (c *StorageController) GetAllStorages(ctx *gin.Context) {
	storages, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(*storages) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Storages Found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"storages": storages})
}
func (c *StorageController) GetAllStoragesByClientId(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	storages, err := c.service.GetAllByClientId(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(*storages) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Storages Found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"storages": storages})
}
func (c *StorageController) GetById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	storage, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if storage.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Storage Found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"storages": storage})
}
func (c *StorageController) CreateStorage(ctx *gin.Context) {
	var dtoAdd StorageDtos.StorageDtoAdd
	if err := ctx.ShouldBind(&dtoAdd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Create(&dtoAdd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
func (c *StorageController) UpdateStorage(ctx *gin.Context) {
	var dtoUpdate StorageDtos.StorageDtoUpdate
	if err := ctx.ShouldBind(&dtoUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Update(&dtoUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
func (c *StorageController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	errDel := c.service.Delete(uint(id))
	if errDel != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
