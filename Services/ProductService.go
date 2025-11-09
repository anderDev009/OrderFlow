package Services

import (
	"orderflow.com/v2/Dto/Productdtos"
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/models"
)

type ProductService struct {
	repository repositories.IProductRepository
	BaseServiceWithClientId[models.Product, Productdtos.DtoAdd, Productdtos.DtoGet, Productdtos.DtoUpdate]
}

// constructor
func NewProductService(repository repositories.IProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
		BaseServiceWithClientId: BaseServiceWithClientId[models.Product, Productdtos.DtoAdd, Productdtos.DtoGet, Productdtos.DtoUpdate]{
			repository: repository,
			BaseService: BaseService[models.Product, Productdtos.DtoAdd, Productdtos.DtoGet, Productdtos.DtoUpdate]{
				repository: repository,
			},
		},
	}
}
