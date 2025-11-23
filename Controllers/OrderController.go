package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"orderflow.com/v2/Dto/OrderDtos"
	"orderflow.com/v2/contracts/services"
)

type OrderController struct {
	service services.IOrderService
}

func NewOrderController(service services.IOrderService) *OrderController {
	return &OrderController{service: service}
}

// obtener las ordenes por cliente
func (controller *OrderController) GetOrdersByClientId(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, errConv := strconv.ParseUint(idStr, 10, 64)
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
		return
	}
	orders, err := controller.service.GetAllByClientId(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}
func (controller *OrderController) GetOrderById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, errConv := strconv.ParseUint(idStr, 10, 64)
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
		return
	}
	orders, err := controller.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

// agregar una orden
func (controller *OrderController) AddOrder(ctx *gin.Context) {
	var order OrderDtos.DtoAdd
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//llamando al servicio
	errAdd := controller.service.Create(&order)
	if errAdd != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errAdd.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

// actualizar una orden
func (controller *OrderController) UpdateOrder(ctx *gin.Context) {
	var order OrderDtos.DtoUpdate
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//llamando al servicio
	errAdd := controller.service.Update(&order)
	if errAdd != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errAdd.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, order)
}
func (controller *OrderController) DeleteOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errDel := controller.service.Delete(uint(id))
	if errDel != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errDel.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// add details
func (controller *OrderController) AddOrderDetail(ctx *gin.Context) {
	var dtoAddDetail OrderDtos.DtoAddDetail
	if err := ctx.ShouldBind(&dtoAddDetail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//calling service
	err := controller.service.AddDetail(&dtoAddDetail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dtoAddDetail)
}

// remove detail
func (controller *OrderController) RemoveDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errDel := controller.service.RemoveDetail(uint(id))
	if errDel != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errDel.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// get order with details
func (controller *OrderController) GetOrderWithDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//calling service
	order, err := controller.service.GetOrderWithDetails(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, order)
}
