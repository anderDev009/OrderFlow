package Services

import (
	"orderflow.com/v2/Dto/OrderDtos"
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/models"
)

type OrderService struct {
	BaseServiceWithClientId[models.Order, OrderDtos.DtoAdd, OrderDtos.DtoGet, OrderDtos.DtoUpdate]
}

func NewOrderService(repository repositories.IOrderRepository) *OrderService {
	base := BaseService[models.Order, OrderDtos.DtoAdd, OrderDtos.DtoGet, OrderDtos.DtoUpdate]{
		repository: repository,
	}
	return &OrderService{
		BaseServiceWithClientId[models.Order, OrderDtos.DtoAdd, OrderDtos.DtoGet, OrderDtos.DtoUpdate]{
			repository:  repository,
			BaseService: &base,
		},
	}
}
