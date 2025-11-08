package Repositories

import "orderflow.com/v2/models"

type ClientRepository struct {
	BaseRepository[models.Client]
}
