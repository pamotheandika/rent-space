package transactions_test

import (
	_spacesMock "RentSpace/businesses/spaces/mocks"
	"RentSpace/businesses/transactions"
	_transactionMock "RentSpace/businesses/transactions/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	transactionRespository _transactionMock.Repository
	transactionService     transactions.Usecase
	transactionDomain      transactions.Domain

	spacesUsecase _spacesMock.Usecase
)

func TestMain(m *testing.M) {
	transactionService = transactions.NewTransactionUsecase(time.Hour*1, &transactionRespository, &spacesUsecase)
	transactionDomain = transactions.Domain{
		ID:              1,
		InvoiceNumber:   "INVSPC0001",
		TransactionDate: "23-09-2021",
		SpaceID:         1,
		SpaceName:       "Depan Bank BRI",
		OwnerID:         1,
		OwnerName:       "Jhon",
		CustomerID:      1,
		CustomerName:    "Legend",
		RentTotalMonth:  3,
		Cost:            300000,
		TotalCost:       900000,
		Status:          1,
	}
}

func TestAddTransaction(t *testing.T) {
	t.Run("Add New Transaction", func(t *testing.T) {
		transactionRespository.On("AddTransaction", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(nil).Once()
		spacesUsecase.On("UpdateStatusSpace", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := transactionService.AddTransaction(context.Background(), &transactionDomain)
		assert.Nil(t, err)
	})

	t.Run("Add New Transaction", func(t *testing.T) {
		transactionRespository.On("AddTransaction", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(nil).Once()
		spacesUsecase.On("UpdateStatusSpace", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := transactionService.AddTransaction(context.Background(), &transactionDomain)
		assert.NotNil(t, err)
	})
}

func TestGetAllTransaction(t *testing.T) {
	t.Run("Get All Transaction", func(t *testing.T) {
		transactionRespository.On("GetAllTransaction", mock.AnythingOfType("context.Context"), mock.AnythingOfType("stirng")).Return([]transactions.Domain{transactionDomain}, nil).Once()

		result, err := transactionService.GetAllTransaction(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All Transaction", func(t *testing.T) {
		transactionRespository.On("GetAllTransaction", mock.AnythingOfType("context.Context"), mock.AnythingOfType("stirng")).Return([]transactions.Domain{transactionDomain}, nil).Once()

		_, err := transactionService.GetAllTransaction(context.Background())
		assert.NotNil(t, err)
	})
}
