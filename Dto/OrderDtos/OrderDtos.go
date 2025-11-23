package OrderDtos

import (
	"orderflow.com/v2/Dto/Customerdtos"
	"orderflow.com/v2/Dto/Productdtos"
	"orderflow.com/v2/Dto/clientdtos"
)

type BaseDto struct {
	CusotmerId uint `json:"customer_id"`
	Status     uint `json:"status"`
	ClientId   uint `json:"client_id"`
}

type DtoGet struct {
	ID uint `json:"id"`
	BaseDto
	Customer Customerdtos.DtoGet
	Client   clientdtos.DtoGet
}
type DtoAdd struct {
	BaseDto
}
type DtoUpdate struct {
	ID uint `json:"id"`
	BaseDto
}

// GetWithDetails
type DtoGetWithDetails struct {
	DtoGet
	Details []DtoGetDetail
}

// dtos to details
type DtoGetDetail struct {
	ID        uint    `json:"id"`
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Productdtos.DtoGet
}

// dto add details
type DtoAddDetail struct {
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
}
