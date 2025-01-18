package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"
)

type GetProductTransactionInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetTransactionInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetTransactionStatusInput struct {
	Status string `json:"status" binding:"required"`
}

type CreateTransactionInput struct {
	ProductID     int    `json:"product_id" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	User          entity.User
}

type TransactionService interface {
	GetTransactionByProductID(input GetProductTransactionInput) ([]entity.Transaction, error)
	GetTransactionByUserID(userID int) ([]entity.Transaction, error)
	GetTransactionByID(input GetTransactionInput) (entity.Transaction, error)
	CreateTransaction(input CreateTransactionInput) (entity.Transaction, error)
	MarkPaid(input GetTransactionInput) (entity.Transaction, error)
	MarkCancel(input GetTransactionInput) (entity.Transaction, error)
	MarkStatus(input GetTransactionInput, inputData GetTransactionStatusInput) (entity.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) GetTransactionByProductID(input GetProductTransactionInput) ([]entity.Transaction, error) {
	transactions, err := s.transactionRepository.GetByProductID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetTransactionByUserID(userID int) ([]entity.Transaction, error) {
	transactions, err := s.transactionRepository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetTransactionByID(input GetTransactionInput) (entity.Transaction, error) {
	transaction, err := s.transactionRepository.FindByID(input.ID)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *transactionService) CreateTransaction(input CreateTransactionInput) (entity.Transaction, error) {
	transaction := entity.Transaction{}
	transaction.UserID = input.User.ID
	transaction.ProductID = input.ProductID
	transaction.Quantity = input.Quantity
	transaction.PaymentMethod = input.PaymentMethod
	// transaction.Total = input.Product.Price * input.Quantity
	transaction.Status = "pending"

	newTransaction, err := s.transactionRepository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *transactionService) MarkPaid(input GetTransactionInput) (entity.Transaction, error) {
	transaction, err := s.transactionRepository.FindByID(input.ID)
	if err != nil {
		return transaction, err
	}

	transaction.Status = "paid"

	markPaid, err := s.transactionRepository.MarkStatus(transaction)
	if err != nil {
		return markPaid, err
	}

	return markPaid, nil
}

func (s *transactionService) MarkCancel(input GetTransactionInput) (entity.Transaction, error) {
	transaction, err := s.transactionRepository.FindByID(input.ID)
	if err != nil {
		return transaction, err
	}

	transaction.Status = "cancel"

	markCancel, err := s.transactionRepository.MarkStatus(transaction)
	if err != nil {
		return markCancel, err
	}

	return markCancel, nil
}

func (s *transactionService) MarkStatus(input GetTransactionInput, inputData GetTransactionStatusInput) (entity.Transaction, error) {
	transaction, err := s.transactionRepository.FindByID(input.ID)
	if err != nil {
		return transaction, err
	}

	transaction.Status = inputData.Status

	markStatus, err := s.transactionRepository.MarkStatus(transaction)
	if err != nil {
		return markStatus, err
	}

	return markStatus, nil
}
