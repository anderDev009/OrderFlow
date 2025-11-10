package repositories

import "orderflow.com/v2/models"

type IStorageRepository interface {
	IBaseRepositoryWithClientId[models.Storage]
}
