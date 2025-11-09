package services

import (
	"orderflow.com/v2/dto/Product"
	"orderflow.com/v2/models"
)

type IProductService interface {
	IBaseService[models.Product, Product.DtoAdd, Product.DtoGet, Product.DtoUpdate]
}
