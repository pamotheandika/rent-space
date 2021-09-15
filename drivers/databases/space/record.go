package space

import "RentSpace/businesses/spaces"

type Space struct {
	IDSpace     int `gorm:"primaryKey"`
	Name        string
	Address     string
	District    string
	City        string
	Province    string
	Email       string
	PhoneNumber string
	Location    string
	IDOwner     int
	Cost        int
	Status      int
}

func (space *Space) toDomain() spaces.Domain {
	return spaces.Domain{
		Name:        space.Name,
		Address:     space.Address,
		District:    space.District,
		City:        space.City,
		Province:    space.Province,
		Email:       space.Email,
		PhoneNumber: space.PhoneNumber,
		IDOwner:     space.IDOwner,
		Cost:        space.Cost,
		Status:      space.Status,
	}
}

func fromDomain(req spaces.Domain) *Space {
	return &Space{
		IDSpace:     req.IDSpace,
		Name:        req.Name,
		Address:     req.Address,
		District:    req.District,
		City:        req.City,
		Province:    req.Province,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		IDOwner:     req.IDOwner,
		Cost:        req.Cost,
		Status:      req.Status,
	}
}

// Method for generate primary
// func createPrimayKey() int {

// }
