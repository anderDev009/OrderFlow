package repositories

type IBaseRepository[T any] interface {
	//Common methods
	GetByID(id uint) (*T, error)
	GetAll() (*[]T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}
