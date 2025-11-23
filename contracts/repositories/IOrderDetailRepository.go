package repositories

import "orderflow.com/v2/models"

type IOrderDetailRepository interface {
	IBaseRepository[models.OrderDetail]
}
