package StorageProductDtos

type BaseDto struct {
	StorageId uint `json:"storage_id"`
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type DtoGet struct {
	ID uint `json:"id"`
	BaseDto
}
type DtoAdd struct {
	BaseDto
}
type DtoUpdate struct {
	ID uint `json:"id"`
	BaseDto
}
type ExtractProductDto struct {
	StorageId int     `json:"storage_id"`
	Quantity  float64 `json:"quantity"`
}
