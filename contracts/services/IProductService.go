package services

import (
	"orderflow.com/v2/Dto/Productdtos"
	"orderflow.com/v2/models"
)

type IProductService interface {
	IBaseServiceWithClientId[models.Product, Productdtos.DtoAdd, Productdtos.DtoGet, Productdtos.DtoUpdate]
}
