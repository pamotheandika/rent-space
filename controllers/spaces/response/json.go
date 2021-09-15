package response

import "RentSpace/businesses/spaces"

type Space struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	District    string `json:"district"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	IDOwner     int    `json:"idowner"`
	Cost        int    `json:"cost"`
}

func FromDomain(spaces spaces.Domain) Space {
	return Space{
		Name:        spaces.Name,
		Address:     spaces.Address,
		District:    spaces.District,
		City:        spaces.City,
		Province:    spaces.Province,
		Email:       spaces.Email,
		PhoneNumber: spaces.PhoneNumber,
		IDOwner:     spaces.IDOwner,
		Cost:        spaces.Cost,
	}
}
