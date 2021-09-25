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

	err = tc.spaceUsecase.UpdateStatusSpace(ctx, transactionDomain.SpaceID)

	if err != nil {
		return err
	}

	return nil
}

func (tc *transactionUsecase) GetAllTransaction(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tc.contextTimeout)
	defer cancel()

	res, err := tc.transactionRepository.GetAllTransaction(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (tc *transactionUsecase) UpdateStatusTransaction(ctx context.Context, IDTransaction int) error {
	ctx, cancel := context.WithTimeout(ctx, tc.contextTimeout)
	defer cancel()

	err := tc.transactionRepository.UpdateStatusTransaction(ctx, IDTransaction)

	if err != nil {
		return err
	}

	return nil
}
