package services

import (
	"orderflow.com/v2/Dto/Customerdtos"
	"orderflow.com/v2/models"
)

type ICustomerService interface {
	IBaseServiceWithClientId[models.Customer, Customerdtos.DtoAdd, Customerdtos.DtoGet, Customerdtos.DtoUpdate]
}
