package Services

import (
	"github.com/jinzhu/copier"
	"orderflow.com/v2/contracts/repositories"
)

type BaseServiceWithClientId[T any, DtoAdd any, DtoGet any, DtoUpdate any] struct {
	repository repositories.IBaseRepositoryWithClientId[T]
	BaseService[T, DtoAdd, DtoGet, DtoUpdate]
}

// constructor
func NewBaseServiceWithClientId[T any, DtoAdd any, DtoGet any, DtoUpdate any](repo repositories.IBaseRepositoryWithClientId[T]) *BaseServiceWithClientId[T, DtoAdd, DtoGet, DtoUpdate] {
	return &BaseServiceWithClientId[T, DtoAdd, DtoGet, DtoUpdate]{
		repository: repo,
		BaseService: BaseService[T, DtoAdd, DtoGet, DtoUpdate]{
			repository: repo,
		},
	}
}

func (s *BaseServiceWithClientId[T, DtoAdd, DtoGet, DtoUpdate]) GetAllByClientId(clientId uint) (*[]DtoGet, error) {
	var dtoGet []DtoGet
	entities, err := s.repository.GetAllByClientId(clientId)
	if err != nil {
		return nil, err
	}
	errCopier := copier.Copy(&dtoGet, entities)
	if errCopier != nil {
		return &[]DtoGet{}, errCopier
	}
	return &dtoGet, nil
}
