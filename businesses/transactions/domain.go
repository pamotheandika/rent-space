package transactions

import (
	"context"
	"time"
)

type Domain struct {
	ID              int
	InvoiceNumber   string
	TransactionDate string
	SpaceID         int
	SpaceName       string
	OwnerID         int
	OwnerName       string
	CustomerID      int
	CustomerName    string
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
	GetAllTransaction(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	AddTransaction(ctx context.Context, transaction *Domain) error
	UpdateStatusTransaction(ctc context.Context, IDTransaction int) error
	GetAllTransaction(ctx context.Context) ([]Domain, error)
}
