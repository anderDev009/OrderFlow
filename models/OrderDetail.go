package models

type OrderDetail struct {
	BaseModel
	OrderId          uint    `gorm:"not null;column:order_id"`
	StorageProductId uint    `gorm:"not null;column:storage_product_id"`
	Quantity         float64 `gorm:"type:decimal(10,2);not null;column:quantity"`
	Price            float64 `gorm:"type:decimal(10,2);not null;column:price"`
	// Navigation properties
	Order          Order          `gorm:"foreignKey:OrderId;references:ID"`
	StorageProduct StorageProduct `gorm:"foreignKey:StorageProductId;references:ID"`
}
