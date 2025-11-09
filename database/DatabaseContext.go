package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

var dbGorm *gorm.DB

func init() {
	dsn := "orderflowuser:orderflowpass@tcp(mysql:3306)/orderflowdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbGorm = db
	errDb := dbGorm.AutoMigrate(&models.Client{},
		&models.Customer{},
		&models.Product{},
		&models.Storage{},
		&models.StorageProduct{},
		&models.Order{},
		&models.OrderDetail{})
	if errDb != nil {
		panic("failed to migrate database")
	}
}

func GetContext() *gorm.DB {
	return dbGorm
}
