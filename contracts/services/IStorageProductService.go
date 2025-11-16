package services

import (
	"orderflow.com/v2/Dto/StorageProductDtos"
	"orderflow.com/v2/models"
)

type IStorageProductService interface {
	IBaseService[models.StorageProduct, StorageProductDtos.DtoAdd,
		StorageProductDtos.DtoGet,
		StorageProductDtos.DtoUpdate]
	ExtractProduct(id int, quantity float64) error
	GetProductsByStorageId(storageID uint) (*[]models.StorageProduct, error)
}
