package Services

import (
	"orderflow.com/v2/Dto/Customerdtos"
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/models"
)

type CustomerService struct {
	BaseServiceWithClientId[models.Customer, Customerdtos.DtoAdd, Customerdtos.DtoGet, Customerdtos.DtoUpdate]
}

// constructor
func NewCustomerService(repository repositories.ICustomerRepository) *CustomerService {
	return &CustomerService{
		BaseServiceWithClientId[models.Customer, Customerdtos.DtoAdd, Customerdtos.DtoGet, Customerdtos.DtoUpdate]{
			repository: repository,
			BaseService: BaseService[models.Customer, Customerdtos.DtoAdd, Customerdtos.DtoGet, Customerdtos.DtoUpdate]{
				repository: repository,
			},
		},
	}
}
