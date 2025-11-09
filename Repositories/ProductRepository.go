package Repositories

import (
	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type ProductRepository struct {
	BaseRepositoryWithClientId[models.Product]
}

// Constructor
func NewProductRepository(ctx *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepositoryWithClientId: BaseRepositoryWithClientId[models.Product]{
			BaseRepository: BaseRepository[models.Product]{
				ctx: ctx,
			},
		},
	}
}
