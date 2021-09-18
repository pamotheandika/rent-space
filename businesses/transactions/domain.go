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
	Status          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type Usecase interface {
	AddTransaction(ctx context.Context, transaction *Domain) error
	UpdateStatusTransaction(ctc context.Context, IDTransaction int) error
}

type Repository interface {
	AddTransaction(ctx context.Context, transaction *Domain) error
	UpdateStatusTransaction(ctc context.Context, IDTransaction int) error
}
