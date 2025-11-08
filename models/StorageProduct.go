package models

type StorageProduct struct {
	BaseModel
	StorageId uint    `gorm:"not null;column:storage_id"`
	ProductId uint    `gorm:"not null;column:product_id"`
	Quantity  float64 `gorm:"type:decimal(10,2);not null;column:quantity"`
	// Navigation properties
	Storage Storage `gorm:"foreignKey:StorageId;references:ID"`
	Product Product `gorm:"foreignKey:ProductId;references:ID"`
}
