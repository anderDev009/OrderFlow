package services

import (
	"orderflow.com/v2/Dto/clientdtos"
	"orderflow.com/v2/models"
)

type IClientService interface {
	IBaseService[models.Client, clientdtos.DtoAdd, clientdtos.DtoGet, clientdtos.DtoUpdate]
}
