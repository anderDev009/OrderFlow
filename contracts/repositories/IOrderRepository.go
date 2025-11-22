package repositories

import "orderflow.com/v2/models"

type IOrderRepository interface {
	IBaseRepositoryWithClientId[models.Order]
	AddDetail(orderDetail models.OrderDetail)
	RemoveDetail(idDetail uint) error
	GetOrderWithDetails(orderId uint) (models.Order, error)
}
