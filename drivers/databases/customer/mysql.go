package customer

import (
	"RentSpace/businesses/customers"
	"context"

	"gorm.io/gorm"
)

type mysqlCustomersRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) customers.Repository {
	return &mysqlCustomersRepository{
		Conn: conn,
	}
}

func (nr *mysqlCustomersRepository) GetByEmail(ctx context.Context, email string) (customers.Domain, error) {
	rec := Customer{}
	result := nr.Conn.Where("email = ?", email).First(&rec)

	if result.Error != nil {
		return customers.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}

func (cr *mysqlCustomersRepository) LoginCustomer(ctx context.Context, email, password string) (customers.Domain, error) {
	rec := Customer{}
	result := cr.Conn.Where("email = ?", email, " password = ?", password).First(&rec)

	if result.Error != nil {
		return customers.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}

func (cr *mysqlCustomersRepository) GetAllCustomer(ctx context.Context) ([]customers.Domain, error) {
	rec := []Customer{}

	resp := cr.Conn.Find(&rec)

	if resp.Error != nil {
		return []customers.Domain{}, resp.Error
	}

	spaceDomains := []customers.Domain{}
	for _, value := range rec {
		spaceDomains = append(spaceDomains, value.toDomain())
	}
	return spaceDomains, nil

}

func (nr *mysqlCustomersRepository) AddCustomer(ctx context.Context, customer *customers.Domain) error {
	rec := fromDomain(*customer)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
