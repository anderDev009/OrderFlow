package Services

import (
	"github.com/jinzhu/copier"
	"orderflow.com/v2/contracts/repositories"
)

type BaseService[T any, DtoAdd any, DtoGet any, DtoUpdate any] struct {
	repository repositories.IBaseRepository[T]
	model      *T
}

// constructor
func NewBaseService[T any, DtoAdd any, DtoGet any, DtoUpdate any](repo repositories.IBaseRepository[T], model *T) *BaseService[T, DtoAdd, DtoGet, DtoUpdate] {
	return &BaseService[T, DtoAdd, DtoGet, DtoUpdate]{
		repository: repo,
		model:      model,
	}
}

// metodos
func (b *BaseService[T, DtoAdd, DtoGet, DtoUpdate]) GetByID(id uint) (*DtoGet, error) {
	var dto DtoGet
	entity, err := b.repository.GetByID(id)
	if err != nil {
		return &dto, err
	}
	//mapping the entity to dto
	copier.Copy(&dto, entity)
	return &dto, nil
}

func (b *BaseService[T, DtoAdd, DtoGet, DtoUpdate]) GetAll() (*[]DtoGet, error) {
	var dto []DtoGet
	entities, err := b.repository.GetAll()
	if err != nil {
		return &[]DtoGet{}, err
	}
	copier.Copy(&dto, entities)
	return &dto, nil
}
func (b *BaseService[T, DtoAdd, DtoGet, DtoUpdate]) Create(DtoAdd *DtoAdd) error {
	var entity T
	copier.Copy(&entity, DtoAdd)
	err := b.repository.Create(&entity)
	if err != nil {
		return err
	}
	return nil
}
func (b *BaseService[T, DtoAdd, DtoGet, DtoUpdate]) Update(DtoUpdate *DtoUpdate) error {
	var entity T
	copier.Copy(&entity, DtoUpdate)
	err := b.repository.Update(&entity)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseService[T, DtoAdd, DtoGet, DtoUpdate]) Delete(id uint) error {
	_, err := b.repository.GetByID(id)
	if err != nil {
		return err
	}
	errDel := b.repository.Delete(id)
	if errDel != nil {
		return err
	}
	return nil
}
