package customers

import (
	"context"
	"time"
)

type Domain struct {
	ID     int
	IdentityNumber string
	Name           string
	Password       string
	BirthOfDate    time.Time
	PlaceOfBirth   string
	Email          string
	PhoneNumber    string
	Address        string
	District       string
	City           string
	Province       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type Usecase interface {
	LoginCustomer(ctx context.Context, email, password string) (string, error)
	AddCustomer(ctx context.Context, customer *Domain) error
	GetAllCustomer(ctx context.Context) ([]Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	// LoginCustomer(ctx context.Context, email, password string) (string, error)
}

type Repository interface {
	AddCustomer(ctx context.Context, customer *Domain) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetAllCustomer(ctx context.Context) ([]Domain, error)
	// LoginCustomer(ctx context.Context, email, password string) (string, error)
	ValidationEmailPassword(ctx context.Context, email, password string) (Domain, error)
}
