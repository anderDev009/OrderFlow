package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func init() {
	dsn := "user:pass@mysql/order_flow?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbGorm = db
}

func GetContext() *gorm.DB {
	return dbGorm
}
