package models

type BaseModel struct {
	ID        uint  `gorm:"primaryKey;autoIncrement;column:id"`
	CreatedAt int64 `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime;column:updated_at"`
}
