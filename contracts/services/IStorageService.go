package services

import (
	"orderflow.com/v2/Dto/StorageDtos"
	"orderflow.com/v2/models"
)

type IStorageService interface {
	IBaseServiceWithClientId[models.Storage, StorageDtos.StorageDtoAdd, StorageDtos.StorageDtoGet, StorageDtos.StorageDtoUpdate]
}
