package owner

import (
	"RentSpace/businesses/owners"
	"time"
)

type Owner struct {
	IDOwner        int `gorm:"primaryKey"`
	IdentityNumber string
	Name           string
	Password       string
	BirthOfDate    string
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

func (req *Owner) toDomain() owners.Domain {
	return owners.Domain{
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

func fromDomain(owner owners.Domain) *Owner {
	return &Owner{
		IdentityNumber: owner.IdentityNumber,
		Name:           owner.Name,
		Password:       owner.Password,
		BirthOfDate:    owner.BirthOfDate,
		PlaceOfBirth:   owner.PlaceOfBirth,
		Email:          owner.Email,
		PhoneNumber:    owner.PhoneNumber,
		Address:        owner.Address,
		District:       owner.District,
		City:           owner.City,
		Province:       owner.Province,
		CreatedAt:      owner.CreatedAt,
		UpdatedAt:      owner.UpdatedAt,
		DeletedAt:      owner.DeletedAt,
	}
}
