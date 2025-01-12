package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"
)

type GetCampaignTransactionInput struct {
	ID int `uri:"id" binding:"required"`
}

type TransactionService interface {
	GetTransactionByProductID(input GetCampaignTransactionInput) ([]entity.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) GetTransactionByProductID(input GetCampaignTransactionInput) ([]entity.Transaction, error) {
	transactions, err := s.transactionRepository.GetByProductID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
