package Repositories

import "gorm.io/gorm"

type BaseRepository[T any] struct {
	//contexto
	ctx *gorm.DB
}

// constructor
func NewBaseRepository[T any](ctx *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{ctx: ctx}
}

// common methods
func (r *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	result := r.ctx.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *BaseRepository[T]) GetAll() (*[]T, error) {
	var entities []T
	result := r.ctx.Find(&entities)
	if result.Error != nil {
		return &[]T{}, result.Error
	}
	return &entities, nil
}
func (r *BaseRepository[T]) Create(entity *T) error {
	result := r.ctx.Create(entity)
	return result.Error
}
func (r *BaseRepository[T]) Update(entity *T) error {
	result := r.ctx.Save(entity)
	return result.Error
}
func (r *BaseRepository[T]) Delete(id uint) error {
	result := r.ctx.Delete(new(T), id)
	return result.Error
}
