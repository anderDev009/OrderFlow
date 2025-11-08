package models

type Client struct {
	BaseModel
	ClientName string `gorm:"type:varchar(100);not null;column:client_name"`
	Rnc        string `gorm:"type:varchar(20);not null;unique;column:rnc"`
	//navigation Property
	Storages  []Storage  `gorm:"foreignKey:ClientId;references:ID"`
	Customers []Customer `gorm:"foreignKey:ClientId;references:ID"`
	Products  []Product  `gorm:"foreignKey:ClientId;references:ID"`
}
