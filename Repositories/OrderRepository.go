package Repositories

import (
	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type OrderRepository struct {
	BaseRepositoryWithClientId[models.Order]
}

func NewOrderRepository(ctx *gorm.DB) *OrderRepository {
	return &OrderRepository{
		BaseRepositoryWithClientId[models.Order]{

			BaseRepository: BaseRepository[models.Order]{
				ctx: ctx,
			},
		},
	}
}
