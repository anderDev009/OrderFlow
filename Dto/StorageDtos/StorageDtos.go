package StorageDtos

type BaseDtoStorage struct {
	Name     string `json:"name" binding:"required"'`
	Location string `json:"location" binding:"required"'`
	ClientId uint   `json:"client_id" binding:"required"'`
}

type StorageDtoGet struct {
	ID uint `json:"id" binding:"required"`
	BaseDtoStorage
}
type StorageDtoAdd struct {
	BaseDtoStorage
}
type StorageDtoUpdate struct {
	ID uint `json:"id" binding:"required"`
	BaseDtoStorage
}
