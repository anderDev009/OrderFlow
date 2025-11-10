package Services

import (
	"github.com/jinzhu/copier"
	"orderflow.com/v2/contracts/repositories"
)

type BaseServiceWithClientId[T any, DtoAdd any, DtoGet any, DtoUpdate any] struct {
	repository repositories.IBaseRepositoryWithClientId[T]
	*BaseService[T, DtoAdd, DtoGet, DtoUpdate]
}

// constructor

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
