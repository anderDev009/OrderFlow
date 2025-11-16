package Repositories

import (
	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type ClientRepository struct {
	BaseRepository[models.Client]
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		BaseRepository: BaseRepository[models.Client]{
			ctx: db,
		},
	}
}
