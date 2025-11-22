package Repositories

import (
	"errors"

	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type OrderRepository struct {
	BaseRepositoryWithClientId[models.Order]
}

func NewOrderRepository(ctx *gorm.DB) *OrderRepository {
	return &OrderRepository{
		BaseRepositoryWithClientId[models.Order]{

			BaseRepository: BaseRepository[models.Order]{
				ctx: ctx,
			},
		},
	}
}

// AddDetail function to add detail to order
func (order *OrderRepository) AddDetail(orderDetail models.OrderDetail) error {
	err := order.ctx.Save(&orderDetail).Error
	if err != nil {
		return err
	}
	return nil
}

// RemoveDetail function to remove details
func (order *OrderRepository) RemoveDetail(idDetail uint) error {
	var orderDetail models.OrderDetail
	err := order.ctx.Find(&orderDetail, idDetail).Error
	if err != nil {
		return err
	}
	if orderDetail.ID == 0 {
		return errors.New("does not exist")
	}
	errDel := order.ctx.Delete(&orderDetail).Error
	if errDel != nil {
		return err
	}
	return nil
}

// GetOrderWithDetails function
func (order *OrderRepository) GetOrderWithDetails(orderId uint) (models.Order, error) {
	var orderDetail models.Order
	//get order with relations
	err := order.ctx.
		Preload("OrderDetail").
		Find(&orderDetail, orderId).Error

	if err != nil {
		return models.Order{}, err
	}
	return orderDetail, nil
}
