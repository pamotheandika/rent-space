package cities

import "context"

type Domain struct {
	ID         int
	WikiDataId string
	City       string
	Name       string
	Country    string
}

type Repository interface {
	GetCities(ctx context.Context) ([]Domain, error)
}

type Usecase interface {
	GetCities(ctx context.Context) ([]Domain, error)	 
}