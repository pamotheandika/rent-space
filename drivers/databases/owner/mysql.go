package owner

import (
	"RentSpace/businesses/owners"
	"context"

	"gorm.io/gorm"
)

type mysqlOwnerRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) owners.Repository {
	return &mysqlOwnerRepository{
		Conn: conn,
	}
}

func (cr *mysqlOwnerRepository) GetOwnerByCity(ctx context.Context, city string) (owners.Domain, error) {
	rec := Owner{}
	result := cr.Conn.Where("city = ?", city).First(&rec)

	if result.Error != nil {
		return owners.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}

func (cr *mysqlOwnerRepository) GetAllOwner(ctx context.Context) ([]owners.Domain, error) {
	rec := []Owner{}

	resp := cr.Conn.Find(&rec)

	if resp.Error != nil {
		return []owners.Domain{}, resp.Error
	}

	ownerDomains := []owners.Domain{}
	for _, value := range rec {
		ownerDomains = append(ownerDomains, value.toDomain())
	}
	return ownerDomains, nil
}

func (nr *mysqlOwnerRepository) GetByEmail(ctx context.Context, username string) (owners.Domain, error) {
	rec := Owner{}
	err := nr.Conn.Where("email = ?", username).First(&rec).Error
	if err != nil {
		return owners.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlOwnerRepository) AddOwner(ctx context.Context, owner *owners.Domain) error {
	rec := fromDomain(*owner)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
