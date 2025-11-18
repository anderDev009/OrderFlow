package services

import (
	"orderflow.com/v2/Dto/OrderDtos"
	"orderflow.com/v2/models"
)

type IOrderService interface {
	IBaseServiceWithClientId[models.Order, OrderDtos.DtoAdd, OrderDtos.DtoGet, OrderDtos.DtoUpdate]
}
