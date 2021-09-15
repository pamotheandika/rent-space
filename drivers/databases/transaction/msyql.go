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
