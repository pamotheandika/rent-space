package transaction

import (
	"RentSpace/businesses/transactions"
	"RentSpace/drivers/databases/customer"
	"RentSpace/drivers/databases/owner"
	"RentSpace/drivers/databases/space"
	"time"
)

type Transaction struct {
	ID              int `gorm:"primaryKey"`
	InvoiceNumber   string
	TransactionDate string
	SpaceID         int
	Space           space.Space
	OwnerID         int
	Owner           owner.Owner
	CustomerID      int
	Customer        customer.Customer
	RentTotalMonth  int
	Cost            int
	TotalCost       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

func fromDomain(transaction *transactions.Domain) *Transaction {
	return &Transaction{
		ID:              transaction.ID,
		InvoiceNumber:   transaction.InvoiceNumber,
		TransactionDate: transaction.TransactionDate,
		SpaceID:         transaction.SpaceID,
		OwnerID:         transaction.OwnerID,
		CustomerID:      transaction.CustomerID,
		RentTotalMonth:  transaction.RentTotalMonth,
		Cost:            transaction.Cost,
		TotalCost:       transaction.TotalCost,
	}
}

func (transaction *Transaction) toDomain() transactions.Domain {
	return transactions.Domain{
		ID:              transaction.ID,
		InvoiceNumber:   transaction.InvoiceNumber,
		TransactionDate: transaction.TransactionDate,
		SpaceID:         transaction.SpaceID,
		SpaceName:       transaction.Space.Name,
		OwnerID:         transaction.OwnerID,
		OwnerName:       transaction.Owner.Name,
		CustomerID:      transaction.CustomerID,
		CustomerName:    transaction.Customer.Name,
		RentTotalMonth:  transaction.RentTotalMonth,
		Cost:            transaction.Cost,
		TotalCost:       transaction.TotalCost,
	}
}
