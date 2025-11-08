package models

type Customer struct {
	BaseModel
	CustomerName string `gorm:"type:varchar(100);not null;column:customer_name"`
	Email        string `gorm:"type:varchar(100);not null;unique;column:email"`
	PhoneNumber  string `gorm:"type:varchar(20);not null;column:phone_number"`
	ClientId     uint   `gorm:"not null;column:client_id"`
	// Navigation property
	Client Client `gorm:"foreignKey:ClientId;references:ID"`
}
