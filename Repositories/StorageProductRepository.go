package Repositories

import (
	"errors"

	"gorm.io/gorm"
	"orderflow.com/v2/models"
)

type StorageProductRepository struct {
	BaseRepository[models.StorageProduct]
}

func NewStorageProductRepository(ctx *gorm.DB) *StorageProductRepository {
	return &StorageProductRepository{
		BaseRepository[models.StorageProduct]{
			ctx: ctx,
		},
	}
}
func (r *StorageProductRepository) Create(entity *models.StorageProduct) error {
	// Verificar si ya existe el StorageProduct
	var register models.StorageProduct

	err := r.ctx.
		Where("storage_id = ?", entity.StorageId).
		First(&register).
		Error

	// Caso: no existe → se crea
	if errors.Is(err, gorm.ErrRecordNotFound) {

		if err := r.ctx.Create(entity).Error; err != nil {
			return err
		}

		return nil
	}

	// Caso: cualquier otro error
	if err != nil {
		return err
	}

	// Si existe → aumentar cantidad
	register.Quantity += entity.Quantity

	// Actualizar SOLO el registro encontrado
	if err := r.ctx.Model(&register).Updates(register).Error; err != nil {
		return err
	}

	return nil
}
func (r *StorageProductRepository) ExtractProduct(id int, quantity float64) error {
	var product models.StorageProduct
	err := r.ctx.First(&product, id).Error
	if err != nil {
		return err
	}
	if product.ID == 0 {
		return errors.New("product not found")
	}

	if (product.Quantity - quantity) < 0 {
		return errors.New("quantity out of range")
	}

	product.Quantity = product.Quantity - quantity

	errUpdate := r.ctx.Save(&product).Error
	if errUpdate != nil {
		return err
	}
	return nil
}
func (r *StorageProductRepository) GetProductsByStorageId(storageID uint) (*[]models.StorageProduct, error) {
	var products []models.StorageProduct
	err := r.ctx.Where("storage_id = ?", storageID).
		Preload("Storage").
		Preload("Product").
		Find(&products).Error
	if err != nil {
		return &[]models.StorageProduct{}, err
	}
	return &products, nil
}
