package repositories

import "orderflow.com/v2/models"

type IProductRepository interface {
	IBaseRepository[models.Product]
}
