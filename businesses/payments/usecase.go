package payments

import (
	"RentSpace/businesses/transactions"
	"context"
	"time"
)

type paymentUsecase struct {
	paymentRepository  Repository
	transactionUsecase transactions.Usecase
	contextTimeout     time.Duration
}

func NewPaymentUsecase(timeout time.Duration, usecase Repository, transactionUsecase transactions.Usecase) Usecase {
	return &paymentUsecase{
		paymentRepository:  usecase,
		contextTimeout:     timeout,
		transactionUsecase: transactionUsecase,
	}
}

func (pc *paymentUsecase) AddPayment(ctx context.Context, paymentDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, pc.contextTimeout)
	defer cancel()

	err := pc.paymentRepository.AddPayment(ctx, paymentDomain)
	if err != nil {
		return err
	}

	err = pc.transactionUsecase.UpdateStatusTransaction(ctx, paymentDomain.IDTransaction)

	if err != nil {
		return err
	}

	return nil
}
