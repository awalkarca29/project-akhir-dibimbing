package helper

import "project-akhir-awal/entity"

type ProductTransactionFormatter struct {
	ID            int    `json:"id"`
	ProductID     int    `json:"product_id"`
	UserID        int    `json:"user_id"`
	Name          string `json:"name"`
	Quantity      int    `json:"quantity"`
	Total         int    `json:"total"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
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
