package services

type IBaseService[T any, DtoAdd any, DtoGet any, DtoUpdate any] interface {
	// Common methods
	GetByID(id uint) (*DtoGet, error)
	GetAll() (*[]DtoGet, error)
	Create(entity *DtoAdd) error
	Update(entity *DtoUpdate) error
	Delete(id uint) error
}
