package models

type Storage struct {
	BaseModel
	Name     string `gorm:"type:varchar(100);not null;column:name"`
	Location string `gorm:"type:varchar(200);not null;column:location"`
	ClientId uint   `gorm:"not null;column:client_id"`
	// Navigation property
	Client         Client           `gorm:"foreignKey:ClientId;references:ID"`
	StorageProduct []StorageProduct `gorm:"foreignKey:StorageId;references:ID"`
}
