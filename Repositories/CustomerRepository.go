package Repositories

import (
	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type CustomerRepository struct {
	BaseRepositoryWithClientId[models.Customer]
}

// constructor
func NewCustomerRepository(ctx *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		BaseRepositoryWithClientId[models.Customer]{
			BaseRepository: BaseRepository[models.Customer]{
				ctx: ctx,
			},
		},
	}
}
