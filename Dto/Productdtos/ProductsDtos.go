package Productdtos

import "orderflow.com/v2/Dto"

type BaseProductDto struct {
	ProductName string  `json:"product_name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Cost        float64 `json:"cost" binding:"required"`
	ClientId    uint    `json:"client_id" binding:"required"`
}
type DtoGet struct {
	BaseDto.BaseDto
	BaseProductDto
}
type DtoAdd struct {
	BaseProductDto
}
type DtoUpdate struct {
	ID uint `json:"id" binding:"required"`
	BaseProductDto
}
