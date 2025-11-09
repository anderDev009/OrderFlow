package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Dto/Customerdtos"
	"orderflow.com/v2/contracts/services"
)

type CustomerController struct {
	service services.ICustomerService
}

func NewCustomerController(service services.ICustomerService) *CustomerController {
	return &CustomerController{
		service: service,
	}
}
func (c *CustomerController) GetCustomers(ctx *gin.Context) {
	customers, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(*customers) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Customers Found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"customers": customers})
}
func (c *CustomerController) GetCustomersWithClientId(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	customers, err := c.service.GetAllByClientId(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(*customers) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Customers Found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"customers": customers})
}
func (c *CustomerController) GetCustomer(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	customer, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if customer.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Customer Found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"customers": customer})
}
func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var add Customerdtos.DtoAdd
	if err := ctx.ShouldBind(&add); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Create(&add)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"customer": add})
}
func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var update Customerdtos.DtoUpdate
	if err := ctx.ShouldBind(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Update(&update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"customer": update})
}
func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
