package repositories

type IBaseRepositoryWithClientId[T any] interface {
	IBaseRepository[T]
	GetAllByClientId(clientId uint) (*[]T, error)
}
