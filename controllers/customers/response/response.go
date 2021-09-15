package response

import "RentSpace/businesses/customers"

type Customer struct {
	IdentityNumber string `json:"identitynumber"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	BirthOfDate    string `json:"birthofdate"`
	PlaceOfBirth   string `json:"placeofbirth"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phonenumber"`
	Address        string `json:"address"`
	District       string `json:"district"`
	City           string `json:"city"`
	Province       string `json:"province"`
}

func FromDomain(domain customers.Domain) Customer {
	return Customer{
		IdentityNumber: domain.IdentityNumber,
		Name:           domain.Name,
		Email:          domain.Email,
	}
}
