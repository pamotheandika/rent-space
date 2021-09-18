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

func (tr *mysqlTransactionRepository) AddTransaction(ctx context.Context, transaction *transactions.Domain) error {
	rec := fromDomain(*transaction)
	result := tr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
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
