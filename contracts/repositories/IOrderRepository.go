package repositories

import "orderflow.com/v2/models"

type IOrderRepository interface {
	IBaseRepositoryWithClientId[models.Order]
}
