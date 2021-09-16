package transactions

import (
	"RentSpace/businesses/spaces"
	"context"
	"time"
)

type transactionUsecase struct {
	transactionRepository Repository
	spaceUsecase          spaces.Usecase
	contextTimeout        time.Duration
}

func NewTransactionUsecase(timeout time.Duration, usecase Repository, spaceUsecase spaces.Usecase) Usecase {
	return &transactionUsecase{
		transactionRepository: usecase,
		contextTimeout:        timeout,
		spaceUsecase:          spaceUsecase,
	}
}

func (tc *transactionUsecase) AddTransaction(ctx context.Context, transactionDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, tc.contextTimeout)
	defer cancel()

	err := tc.transactionRepository.AddTransaction(ctx, transactionDomain)
	if err != nil {
		return err
	}

	err = tc.spaceUsecase.UpdateStatusSpace(ctx, transactionDomain.IDSpace)

	if err != nil {
		return err
	}

	return nil
}
