package Repositories

import (
	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type ProductRepository struct {
	BaseRepository[models.Product]
}

// Constructor
func NewProductRepository(ctx *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepository: BaseRepository[models.Product]{
			ctx: ctx,
		},
	}
}
