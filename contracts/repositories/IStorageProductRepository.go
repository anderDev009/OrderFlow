package repositories

import "orderflow.com/v2/models"

type IStorageProductRepository interface {
	IBaseRepository[models.StorageProduct]
	ExtractProduct(id int, quantity float64) error
	GetProductsByStorageId(storageID uint) (*[]models.StorageProduct, error)
}
