package OrderDtos

import (
	"orderflow.com/v2/Dto/Customerdtos"
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
