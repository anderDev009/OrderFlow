package Customerdtos

type DtoBase struct {
	CustomerName string `json:"customer_name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
	ClientId     uint   `json:"client_id" binding:"required"`
}

type DtoGet struct {
	ID uint `json:"id"`
	DtoBase
}

type DtoAdd struct {
	DtoBase
}
type DtoUpdate struct {
	ID uint `json:"id"`
	DtoBase
}
