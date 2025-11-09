package repositories

import "orderflow.com/v2/models"

type ICustomerRepository interface {
	IBaseRepositoryWithClientId[models.Customer]
}
