package services

import (
	"orderflow.com/v2/Dto/OrderDtos"
	"orderflow.com/v2/models"
)

type IOrderService interface {
	IBaseServiceWithClientId[models.Order, OrderDtos.DtoAdd, OrderDtos.DtoGet, OrderDtos.DtoUpdate]
	AddDetail(orderDetail *OrderDtos.DtoAddDetail) error
	RemoveDetail(idDetail uint) error
	GetOrderWithDetails(orderId uint) (OrderDtos.DtoGetWithDetails, error)
}
