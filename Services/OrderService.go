package Services

import (
	"github.com/jinzhu/copier"
	"orderflow.com/v2/Dto/OrderDtos"
	"orderflow.com/v2/contracts/repositories"
	"orderflow.com/v2/models"
)

type OrderService struct {
	BaseServiceWithClientId[models.Order, OrderDtos.DtoAdd, OrderDtos.DtoGet, OrderDtos.DtoUpdate]
	repository repositories.IOrderRepository
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
		repository,
	}
}
func (service *OrderService) AddDetail(orderDetail *OrderDtos.DtoAddDetail) error {
	var orderModel models.OrderDetail
	//mapping the dto to model
	errCopy := copier.Copy(&orderModel, orderDetail)
	if errCopy != nil {
		return errCopy
	}
	//calling repository
	errAdd := service.repository.AddDetail(&orderModel)
	if errAdd != nil {
		return errAdd
	}
	return nil
}

func (service *OrderService) RemoveDetail(idDetail uint) error {
	errDel := service.repository.RemoveDetail(idDetail)
	if errDel != nil {
		return errDel
	}
	return nil
}
func (service *OrderService) GetOrderWithDetails(orderId uint) (OrderDtos.DtoGetWithDetails, error) {
	var orderDtoGet OrderDtos.DtoGetWithDetails
	//getting model to map to OrderDto
	orderModel, err := service.repository.GetOrderWithDetails(orderId)
	if err != nil {
		return OrderDtos.DtoGetWithDetails{}, err
	}
	err = copier.Copy(&orderModel, &orderDtoGet)
	if err != nil {
		return OrderDtos.DtoGetWithDetails{}, err
	}
	return orderDtoGet, nil
}
