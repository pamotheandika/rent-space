package response

import "RentSpace/businesses/owners"

type Owner struct {
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

func FromDomain(domain owners.Domain) Owner {
	return Owner{
		IdentityNumber: domain.IdentityNumber,
		Name:           domain.Name,
		Email:          domain.Email,
		BirthOfDate: domain.BirthOfDate,
		PlaceOfBirth: domain.PlaceOfBirth,
		PhoneNumber: domain.PhoneNumber,
		Address: domain.Address,
		District: domain.District,
		City: domain.City,
		Province: domain.Province,
	}
}
