package repositories

import "orderflow.com/v2/models"

type IProductRepository interface {
	IBaseRepositoryWithClientId[models.Product]
}
