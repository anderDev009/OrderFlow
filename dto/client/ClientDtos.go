package client

import "orderflow.com/v2/dto"

type baseClientDto struct {
	ClientName string `json:"client_name"`
	Rnc        string `json:"rnc"`
}
type DtoGet struct {
	dto.BaseDto
	baseClientDto
}
type DtoUpdate struct {
	ID uint `json:"id"`
	baseClientDto
}
type DtoAdd struct {
	baseClientDto
}
