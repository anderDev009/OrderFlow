package repositories

import "orderflow.com/v2/models"

type IClientRepository interface {
	IBaseRepository[models.Client]
}
