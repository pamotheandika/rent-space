package transaction

import (
	"RentSpace/businesses/transactions"
	"context"

	"gorm.io/gorm"
)

type mysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) transactions.Repository {
	return &mysqlTransactionRepository{
		Conn: conn,
	}
}

func (nr *mysqlTransactionRepository) GetAllTransaction(ctx context.Context) ([]transactions.Domain, error) {
	rec := []Transaction{}

	err := nr.Conn.Preload("Space").Preload("Owner").Preload("Customer").Find(&rec).Error

	if err != nil {
		return []transactions.Domain{}, err
	}

	var domainNews []transactions.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, nil
}

func (tr *mysqlTransactionRepository) AddTransaction(ctx context.Context, transaction *transactions.Domain) error {
	rec := fromDomain(transaction)
	result := tr.Conn.Create(rec)

	if result.Error != nil {
		return result.Error
	}

	err := tr.Conn.Preload("Space").First(&rec, rec.ID).Error
	if err != nil {
		return err
	}

	err = tr.Conn.Preload("Owner").First(&rec, rec.ID).Error
	if err != nil {
		return err
	}

	err = tr.Conn.Preload("Customer").First(&rec, rec.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (sr *mysqlTransactionRepository) UpdateStatusTransaction(ctx context.Context, IDTransaction int) error {
	rec := Transaction{}
	resp := sr.Conn.Where("id_transaction = ?", IDTransaction).Find(&rec).Update("status", 1)

	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
