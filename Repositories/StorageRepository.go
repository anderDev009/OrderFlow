package Repositories

import (
	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type StorageRepository struct {
	BaseRepositoryWithClientId[models.Storage]
}

// constructor
func NewStorageRepository(ctx *gorm.DB) *StorageRepository {
	return &StorageRepository{
		BaseRepositoryWithClientId[models.Storage]{
			BaseRepository: BaseRepository[models.Storage]{
				ctx: ctx,
			},
		},
	}
}
