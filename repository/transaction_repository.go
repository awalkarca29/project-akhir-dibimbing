package repository

import (
	"project-akhir-awal/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetByProductID(productID int) ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetByProductID(productID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Where("product_id = ?", productID).Order("id desc").Find(&transactions).Error
	// err := r.db.Preload("User").Preload("Product").Where("product_id = ?", productID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
