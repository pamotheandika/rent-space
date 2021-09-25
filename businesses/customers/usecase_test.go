package customers_test

import (
	"RentSpace/app/middleware"
	"RentSpace/businesses/customers"
	_customerMock "RentSpace/businesses/customers/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	customerRepository _customerMock.Repository
	customerService    customers.Usecase
	customerDomain     customers.Domain
	jwtAuth            middleware.ConfigJwt
)

func TestMain(m *testing.M) {
	customerService = customers.NewCustomerUsace(time.Hour*1, &customerRepository, &jwtAuth)
	customerDomain = customers.Domain{
		ID:             1,
		IdentityNumber: "12345678910",
		Name:           "Name Test",
		Password:       "Name1234",
		Email:          "test@example.com",
		BirthOfDate:    time.Now(),
		PlaceOfBirth:   "Palopo",
		PhoneNumber:    "12345678",
		Address:        "Jln Bangau",
		District:       "Tamalate",
		City:           "Makassar",
		Province:       "Sulawesi Selatan",
	}
}

func TestGetAllCustomer(t *testing.T) {
	t.Run("Get All Data Customer", func(t *testing.T) {
		customerRepository.On("GetAllCustomer", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return([]customers.Domain{customerDomain}, nil).Once()

		result, err := customerService.GetAllCustomer(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))

	})

	t.Run("Get All Data Customer", func(t *testing.T) {
		customerRepository.On("GetAllCustomer", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return([]customers.Domain{customerDomain}, nil).Once()

		_, err := customerService.GetAllCustomer(context.Background())
		assert.NotNil(t, err)
	})
}

func TestLoginCustomer(t *testing.T) {
	t.Run("Login Customer", func(t *testing.T) {
		customerRepository.On("LoginCustomer", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return(customerDomain, nil).Once()

		result, err := customerService.LoginCustomer(context.Background(), customerDomain.Name, customerDomain.Password)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))

	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get Data By ID", func(t *testing.T) {
		customerRepository.On("GetByID", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return(customerDomain, nil).Once()

		result, err := customerService.GetByID(context.Background(), customerDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)

	})
}

func TestGetByEmail(t *testing.T) {
	t.Run("Get Data By Email", func(t *testing.T) {
		customerRepository.On("GetByEmail", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return(customerDomain, nil).Once()

		result, err := customerService.GetByEmail(context.Background(), customerDomain.Email)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)

	})
}

func TestAddCustomer(t *testing.T) {
	t.Run("Add New Customer", func(t *testing.T) {
		customerRepository.On("AddCustomer", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return(nil).Once()

		err := customerService.AddCustomer(context.Background(), &customerDomain)
		assert.Nil(t, err)
	})
}
