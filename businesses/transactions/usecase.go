package transactions

import (
	"context"
	"time"
)

type transactionUsecase struct {
	transactionRepository Repository
	contextTimeout        time.Duration
}

func NewTransactionUsecase(timeout time.Duration, usecase Repository) Usecase {
	return &transactionUsecase{
		transactionRepository: usecase,
		contextTimeout:        timeout,
	}
}

func (tc *transactionUsecase) AddTransaction(ctx context.Context, transactionDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, tc.contextTimeout)
	defer cancel()

	err := tc.transactionRepository.AddTransaction(ctx, transactionDomain)
	if err != nil {
		return err
	}

	return nil
}
