package services

type IBaseServiceWithClientId[T any, DtoAdd any, DtoGet any, DtoUpdate any] interface {
	IBaseService[T, DtoAdd, DtoGet, DtoUpdate]
	GetAllByClientId(clientId uint) (*[]DtoGet, error)
}
