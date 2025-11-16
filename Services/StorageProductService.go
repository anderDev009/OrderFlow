package Services

import (
	"orderflow.com/v2/Dto/StorageProductDtos"
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/models"
)

type StorageProductService struct {
	repository repositories.IStorageProductRepository
	BaseService[models.StorageProduct, StorageProductDtos.DtoAdd,
		StorageProductDtos.DtoGet,
		StorageProductDtos.DtoUpdate]
}

func NewStorageProductService(repository repositories.IStorageProductRepository) *StorageProductService {
	return &StorageProductService{
		repository,
		BaseService[models.StorageProduct, StorageProductDtos.DtoAdd,
			StorageProductDtos.DtoGet,
			StorageProductDtos.DtoUpdate]{
			repository: repository,
		},
	}
}
func (r *StorageProductService) ExtractProduct(id int, quantity float64) error {
	err := r.repository.ExtractProduct(id, quantity)
	return err
}
func (r *StorageProductService) GetProductsByStorageId(storageID uint) (*[]models.StorageProduct, error) {
	products, err := r.repository.GetProductsByStorageId(storageID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
