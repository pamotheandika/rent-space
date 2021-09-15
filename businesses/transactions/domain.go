package transactions

import (
	"context"
	"time"
)

type Domain struct {
	IDTransaction   int
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

type Usecase interface {
	AddTransaction(ctx context.Context, transaction *Domain) error
}

type Repository interface {
	AddTransaction(ctx context.Context, transaction *Domain) error
}
