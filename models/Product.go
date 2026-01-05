package models

type Product struct {
	BaseModel
	ProductName string  `gorm:"type:varchar(100);not null;column:product_name"`
	Price       float64 `gorm:"type:decimal(10,2);not null;column:price"`
	Cost        float64 `gorm:"type:decimal(10,2);not null;column:cost"`
	ClientId    uint    `gorm:"not null;column:client_id"`
	//navegation property
	Client Client `gorm:"foreignKey:ClientId;references:ID"`
	//Nav property
	StorageProducts []StorageProduct `gorm:"foreignKey:ProductId;references:ID"`
	// OrderDetails    []OrderDetail    `gorm:"foreignKey:ProductId;references:ID"`
}
