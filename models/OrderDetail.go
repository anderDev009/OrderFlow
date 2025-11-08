package models

type OrderDetail struct {
	BaseModel
	OrderId   uint    `gorm:"not null;column:order_id"`
	ProductId uint    `gorm:"not null;column:product_id"`
	Quantity  float64 `gorm:"type:decimal(10,2);not null;column:quantity"`
	Price     float64 `gorm:"type:decimal(10,2);not null;column:price"`
	// Navigation properties
	Order   Order   `gorm:"foreignKey:OrderId;references:ID"`
	Product Product `gorm:"foreignKey:ProductId;references:ID"`
}
