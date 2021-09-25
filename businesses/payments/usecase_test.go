package payments_test

import (
	"RentSpace/businesses/payments"
	_paymentMock "RentSpace/businesses/payments/mocks"
	_transactionMock "RentSpace/businesses/transactions/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	transactionUsecase _transactionMock.Usecase
	paymentRepository  _paymentMock.Repository
	paymentService     payments.Usecase
	paymentDomain      payments.Domain
)

func TestMain(m *testing.M) {
	paymentService = payments.NewPaymentUsecase(time.Hour*1, &paymentRepository, &transactionUsecase)
	paymentDomain = payments.Domain{
		ID:            1,
		IDTransaction: 1,
		InvoiceNumber: "INVSPC0001",
		PaymentAmount: 3500000,
	}
}

func TestAddPayment(t *testing.T) {
	t.Run("Add Payment", func(t *testing.T) {
		paymentRepository.On("AddPayment", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(nil).Once()
		transactionUsecase.On("UpdateStatusTransaction", mock.Anything, mock.AnythingOfType("int")).Once()

		err := paymentService.AddPayment(context.Background(), &paymentDomain)
		assert.Nil(t, err)
	})

	t.Run("Get All Invalid", func(t *testing.T) {
		paymentRepository.On("AddPayment", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(nil).Once()
		transactionUsecase.On("UpdateStatusTransaction", mock.Anything, mock.AnythingOfType("int")).Once()

		err := paymentService.AddPayment(context.Background(), &paymentDomain)
		assert.NotNil(t, err)
	})
}
