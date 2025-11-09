package Services

import (
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/dto/Product"
	"orderflow.com/v2/models"
)

type ProductService struct {
	repository repositories.IProductRepository
	BaseService[models.Product, Product.DtoAdd, Product.DtoGet, Product.DtoUpdate]
}

// constructor
func NewProductService(repository repositories.IProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
		BaseService: BaseService[models.Product, Product.DtoAdd, Product.DtoGet, Product.DtoUpdate]{
			repository: repository,
		},
	}
}
