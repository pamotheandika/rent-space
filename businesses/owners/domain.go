package owners

import (
	"context"
	"time"
)

type Domain struct {
	ID             int
	IdentityNumber string
	Name           string
	BirthOfDate    string
	PlaceOfBirth   string
	Email          string
	PhoneNumber    string
	Address        string
	District       string
	City           string
	Province       string
	Password       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type Usecase interface {
	AddOwner(ctx context.Context, owner *Domain) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetAllOwner(ctx context.Context) ([]Domain, error)
	GetOwnerByCity(ctx context.Context, city string) (Domain, error)
}

type Repository interface {
	AddOwner(ctx context.Context, owner *Domain) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetAllOwner(ctx context.Context) ([]Domain, error)
	GetOwnerByCity(ctx context.Context, city string) (Domain, error)
}
