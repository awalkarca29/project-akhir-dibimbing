package repository

import (
	"project-akhir-awal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindByID(ID int) (entity.Product, error)
	Save(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(product entity.Product) (entity.Product, error)
	UploadImage(productImage entity.ProductImage) (entity.ProductImage, error)
	MarkAllImagesAsNonPrimary(productID int) (bool, error)
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

func (r *productRepository) FindByID(ID int) (entity.Product, error) {
	var product entity.Product

	err := r.db.Preload("ProductImages").Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) Save(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Update(product entity.Product) (entity.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Delete(product entity.Product) (entity.Product, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) UploadImage(productImage entity.ProductImage) (entity.ProductImage, error) {
	err := r.db.Create(&productImage).Error
	if err != nil {
		return productImage, err
	}

	return productImage, nil
}

func (r *productRepository) MarkAllImagesAsNonPrimary(productID int) (bool, error) {
	err := r.db.Model(&entity.ProductImage{}).Where("product_id = ?", productID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
