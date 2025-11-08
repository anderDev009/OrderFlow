package Services

import (
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/contracts/services"
)

type ClientService struct {
	service    services.IClientService
	repository repositories.IClientRepository
}

// constructor
func NewClientService(repo repositories.IClientRepository, service services.IClientService) *ClientService {
	return &ClientService{
		service:    service,
		repository: repo,
	}
}
