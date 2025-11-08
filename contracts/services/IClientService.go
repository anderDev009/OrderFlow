package services

import (
	"orderflow.com/v2/dto/client"
	"orderflow.com/v2/models"
)

type IClientService interface {
	IBaseService[models.Client, client.DtoAdd, client.DtoGet, client.DtoUpdate]
}
