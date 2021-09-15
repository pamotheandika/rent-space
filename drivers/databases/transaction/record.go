package transaction

import (
	"RentSpace/businesses/transactions"
	"time"
)

type Transaction struct {
	IDTransaction   int `gorm:"primaryKey"`
	InvoiceNumber   string
	TransactionDate string
	IDSpace         int
	IDOwner         int
	IDCustomer      int
	RentTotalMonth  int
	Cost            int
	TotalCost       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

func (transaction *Transaction) toDomain() transactions.Domain {
	return transactions.Domain{
		IDTransaction:   transaction.IDTransaction,
		InvoiceNumber:   transaction.InvoiceNumber,
		TransactionDate: transaction.TransactionDate,
		IDSpace:         transaction.IDSpace,
		IDOwner:         transaction.IDOwner,
		IDCustomer:      transaction.IDCustomer,
		RentTotalMonth:  transaction.RentTotalMonth,
		Cost:            transaction.Cost,
		TotalCost:       transaction.TotalCost,
	}
}

func fromDomain(transaction transactions.Domain) *Transaction {
	return &Transaction{
		IDTransaction:   transaction.IDTransaction,
		InvoiceNumber:   transaction.InvoiceNumber,
		TransactionDate: transaction.TransactionDate,
		IDSpace:         transaction.IDSpace,
		IDOwner:         transaction.IDOwner,
		IDCustomer:      transaction.IDCustomer,
		RentTotalMonth:  transaction.RentTotalMonth,
		Cost:            transaction.Cost,
		TotalCost:       transaction.TotalCost,
	}
}
