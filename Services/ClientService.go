package Services

import (
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/dto/client"
	"orderflow.com/v2/models"
)

type ClientService struct {
	repository repositories.IClientRepository
	BaseService[models.Client, client.DtoAdd, client.DtoGet, client.DtoUpdate]
}

// constructor
func NewClientService(repo repositories.IClientRepository) *ClientService {
	return &ClientService{
		repository: repo,
		BaseService: BaseService[models.Client, client.DtoAdd, client.DtoGet, client.DtoUpdate]{
			repository: repo,
		},
	}
}
