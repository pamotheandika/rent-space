package response

import "RentSpace/businesses/transactions"

type Transaction struct {
	ID              int
	InvoiceNumber   string
	TransactionDate string
	SpaceName       string
	OwnerName       string
	CustomerName    string
	RentTotalMonth  int
	TotalCost       int
	Status          int
}

func FromDomain(transaction transactions.Domain) Transaction {
	return Transaction{
		ID:              transaction.ID,
		InvoiceNumber:   transaction.InvoiceNumber,
		TransactionDate: transaction.TransactionDate,
		SpaceName:       transaction.SpaceName,
		OwnerName:       transaction.OwnerName,
		CustomerName:    transaction.CustomerName,
		RentTotalMonth:  transaction.RentTotalMonth,
		TotalCost:       transaction.TotalCost,
		Status:          transaction.Status,
	}
}
