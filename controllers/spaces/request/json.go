package request

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

func (req *Space) ToDomain() *spaces.Domain {
	return &spaces.Domain{
		Name:        req.Name,
		Address:     req.Address,
		District:    req.District,
		City:        req.City,
		Province:    req.Province,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		IDOwner:     req.IDOwner,
		Cost:        req.Cost,
	}
}