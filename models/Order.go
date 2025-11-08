package models

type Order struct {
	BaseModel
	CustomerId uint   `gorm:"not null;column:customer_id"`
	Status     string `gorm:"type:varchar(50);not null;column:status"`
	ClientId   uint   `gorm:"not null;column:client_id"`
	// Navigation properties
	Customer Customer      `gorm:"foreignKey:CustomerId;references:ID"`
	Client   Client        `gorm:"foreignKey:ClientId;references:ID"`
	Details  []OrderDetail `gorm:"foreignKey:OrderId;references:ID"`
}
