package repository

import (
	"project-akhir-awal/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetByProductID(productID int) ([]entity.Transaction, error)
	GetByUserID(UserID int) ([]entity.Transaction, error)
	FindByID(ID int) (entity.Transaction, error)
	Save(transaction entity.Transaction) (entity.Transaction, error)
	MarkStatus(transaction entity.Transaction) (entity.Transaction, error)
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

func (r *transactionRepository) GetByUserID(UserID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Preload("Product.ProductImages", "product_images.is_primary = 1").Preload("User").Where("user_id = ?", UserID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) FindByID(ID int) (entity.Transaction, error) {
	var transaction entity.Transaction

	err := r.db.Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepository) Save(transaction entity.Transaction) (entity.Transaction, error) {
	var product entity.Product
	r.db.First(&product, transaction.ProductID)

	transaction.Total = product.Price * transaction.Quantity
	err := r.db.Preload("Product").Preload("User").Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) MarkStatus(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
