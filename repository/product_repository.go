package repository

import (
	"project-akhir-awal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Preload("ProductImages", "product_images.is_primary = 1").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}
