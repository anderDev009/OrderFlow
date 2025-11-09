package Services

import (
	"orderflow.com/v2/Dto/clientdtos"
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/models"
)

type ClientService struct {
	repository repositories.IClientRepository
	BaseService[models.Client, clientdtos.DtoAdd, clientdtos.DtoGet, clientdtos.DtoUpdate]
}

// constructor
func NewClientService(repo repositories.IClientRepository) *ClientService {
	return &ClientService{
		repository: repo,
		BaseService: BaseService[models.Client, clientdtos.DtoAdd, clientdtos.DtoGet, clientdtos.DtoUpdate]{
			repository: repo,
		},
	}
}
