package response

import (
	"RentSpace/businesses/customers"
	"time"
)

type Customer struct {
	IdentityNumber string    `json:"identitynumber"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	BirthOfDate    time.Time `json:"birthofdate"`
	PlaceOfBirth   string    `json:"placeofbirth"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phonenumber"`
	Address        string    `json:"address"`
	District       string    `json:"district"`
	City           string    `json:"city"`
	Province       string    `json:"province"`
	Token          string    `json:"token"`
}

func FromDomain(domain customers.Domain) Customer {
	return Customer{
		IdentityNumber: domain.IdentityNumber,
		Name:           domain.Name,
		Email:          domain.Email,
		BirthOfDate:    domain.BirthOfDate,
		PlaceOfBirth:   domain.PlaceOfBirth,
		Address:        domain.Address,
		District:       domain.District,
		City:           domain.City,
		Province:       domain.Province,
	}
}
