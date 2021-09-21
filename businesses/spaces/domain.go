package spaces

import "context"

type Domain struct {
	ID     int
	Name        string
	Address     string
	District    string
	City        string
	Province    string
	Email       string
	PhoneNumber string
	IDOwner     int
	Cost        int
	Status      int
}

type Usecase interface {
	AddSpace(ctx context.Context, space *Domain) error
	GetAllSpace(ctx context.Context) ([]Domain, error)
	UpdateStatusSpace(ctx context.Context, IDSpace int) error
}

type Repository interface {
	AddSpace(ctx context.Context, space *Domain) error
	GetAllSpace(ctx context.Context) ([]Domain, error)
	UpdateStatusSpace(ctx context.Context, IDSpace int) error
}
