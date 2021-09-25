package payments

import (
	"context"
	"time"
)

type Domain struct {
	ID            int
	IDTransaction int
	InvoiceNumber string
	PaymentAmount int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type Usecase interface {
	AddPayment(context context.Context, paymentDomain *Domain) error
}

type Repository interface {
	AddPayment(context context.Context, paymentDomain *Domain) error
}
