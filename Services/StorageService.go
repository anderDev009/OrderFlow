package Services

import (
	"orderflow.com/v2/Dto/StorageDtos"
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/models"
)

type StorageService struct {
	BaseServiceWithClientId[models.Storage, StorageDtos.StorageDtoAdd, StorageDtos.StorageDtoGet, StorageDtos.StorageDtoUpdate]
}

// constructor
func NewStorageService(repository repositories.IStorageRepository) *StorageService {
	return &StorageService{
		BaseServiceWithClientId: BaseServiceWithClientId[models.Storage, StorageDtos.StorageDtoAdd, StorageDtos.StorageDtoGet, StorageDtos.StorageDtoUpdate]{
			repository: repository,
			BaseService: &BaseService[models.Storage, StorageDtos.StorageDtoAdd, StorageDtos.StorageDtoGet, StorageDtos.StorageDtoUpdate]{
				repository: repository,
			},
		},
	}
}
