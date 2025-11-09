package Repositories

import "gorm.io/gorm"

type BaseRepositoryWithClientId[T any] struct {
	BaseRepository[T]
}

// constructor
func NewBaseRepositoryWithClientId[T any](ctx *gorm.DB) *BaseRepositoryWithClientId[T] {
	return &BaseRepositoryWithClientId[T]{
		BaseRepository: BaseRepository[T]{
			ctx: ctx,
		},
	}
}

// methods
func (c *BaseRepositoryWithClientId[T]) GetAllByClientId(clientId uint) (*[]T, error) {
	var entities []T
	err := c.ctx.Where("client_id = ?", clientId).Find(&entities).Error
	if err != nil {
		return &[]T{}, err
	}
	return &entities, nil
}
