package customer

import (
	"RentSpace/businesses/customers"
	"time"
)

type Customer struct {
	IDCostumer     int `gorm:"primaryKey"`
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

func (req *Customer) toDomain() customers.Domain {
	return customers.Domain{
		IdentityNumber: req.IdentityNumber,
		Name:           req.Name,
		Password:       req.Password,
		BirthOfDate:    req.BirthOfDate,
		PlaceOfBirth:   req.PlaceOfBirth,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		Address:        req.Address,
		District:       req.District,
		City:           req.City,
		Province:       req.Province,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
		DeletedAt:      req.DeletedAt,
	}
}

func fromDomain(customer customers.Domain) *Customer {
	return &Customer{
		IdentityNumber: customer.IdentityNumber,
		Name:           customer.Name,
		Password:       customer.Password,
		BirthOfDate:    customer.BirthOfDate,
		PlaceOfBirth:   customer.PlaceOfBirth,
		Email:          customer.Email,
		PhoneNumber:    customer.PhoneNumber,
		Address:        customer.Address,
		District:       customer.District,
		City:           customer.City,
		Province:       customer.Province,
		CreatedAt:      customer.CreatedAt,
		UpdatedAt:      customer.UpdatedAt,
		DeletedAt:      customer.DeletedAt,
	}
}