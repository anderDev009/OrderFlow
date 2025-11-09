package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"orderflow.com/v2/contracts/services"
	"orderflow.com/v2/dto/Product"
)

type ProductController struct {
	service services.IProductService
}

func NewProductController(service services.IProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) ListProducts(ctx *gin.Context) {
	products, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(*products) == 0 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, products)
}
func (c *ProductController) AddProduct(ctx *gin.Context) {
	var dtoAdd Product.DtoAdd
	if err := ctx.ShouldBind(&dtoAdd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//llamada al servicio
	err := c.service.Create(&dtoAdd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var dtoUpdate Product.DtoUpdate
	if err := ctx.ShouldBind(&dtoUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//llamada al servicio
	err := c.service.Update(&dtoUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errDel := c.service.Delete(uint(id))
	if errDel != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
func (c *ProductController) GetProductById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
func (c *ProductController) GetProductListWithClientId(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := c.service.GetAllByClientId(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(*products) == 0 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, products)
}
