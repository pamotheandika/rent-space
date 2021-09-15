package request

import (
	"RentSpace/businesses/owners"
)

type Owners struct {
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

func (req *Owners) ToDomain() *owners.Domain {
	return &owners.Domain{
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
	}
}
