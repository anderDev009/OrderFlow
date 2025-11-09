package Customerdtos

type dtoBase struct {
	CustomerName string `json:"customer_name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
	ClientId     uint   `json:"client_id" binding:"required"`
}

type DtoGet struct {
	ID uint `json:"id"`
	dtoBase
}

type DtoAdd struct {
	dtoBase
}
type DtoUpdate struct {
	ID uint `json:"id"`
	dtoBase
}
