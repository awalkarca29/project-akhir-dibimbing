package helper

import (
	"project-akhir-awal/entity"
	"time"
)

type ProductTransactionFormatter struct {
	ID            int       `json:"id"`
	ProductID     int       `json:"product_id"`
	UserID        int       `json:"user_id"`
	Name          string    `json:"name"`
	Quantity      int       `json:"quantity"`
	Total         int       `json:"total"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

func FormatProductTransaction(transaction entity.Transaction) ProductTransactionFormatter {
	formatter := ProductTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.ProductID = transaction.ProductID
	formatter.UserID = transaction.UserID
	formatter.Name = transaction.User.Name
	formatter.Quantity = transaction.Quantity
	formatter.Total = transaction.Total
	formatter.PaymentMethod = transaction.PaymentMethod
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatProductTransactions(transactions []entity.Transaction) []ProductTransactionFormatter {
	if len(transactions) == 0 {
		return []ProductTransactionFormatter{}
	}
	var transactionsFormatter []ProductTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatProductTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID            int                       `json:"id"`
	ProductID     int                       `json:"product_id"`
	UserID        int                       `json:"user_id"`
	Name          string                    `json:"name"`
	Quantity      int                       `json:"quantity"`
	Total         int                       `json:"total"`
	PaymentMethod string                    `json:"payment_method"`
	Status        string                    `json:"status"`
	CreatedAt     time.Time                 `json:"created_at"`
	Product       ProductUserTransFormatter `json:"campaign"`
}

type ProductUserTransFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction entity.Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.ProductID = transaction.ProductID
	formatter.UserID = transaction.UserID
	formatter.Name = transaction.User.Name
	formatter.Quantity = transaction.Quantity
	formatter.Total = transaction.Total
	formatter.PaymentMethod = transaction.PaymentMethod
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	productFormatter := ProductUserTransFormatter{}
	productFormatter.Name = transaction.Product.Name

	if len(transaction.Product.ProductImages) > 0 {
		productFormatter.ImageURL = transaction.Product.ProductImages[0].FileName
	}

	formatter.Product = productFormatter

	return formatter
}

func FormatUserTransactions(transactions []entity.Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}
	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
