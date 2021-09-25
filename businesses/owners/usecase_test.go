package owners_test

import (
	"RentSpace/app/middleware"
	"RentSpace/businesses/owners"
	_ownerMocks "RentSpace/businesses/owners/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	ownerRepository _ownerMocks.Repository
	ownerService    owners.Usecase
	ownerDomain     owners.Domain
	jwtAuth         middleware.ConfigJwt
)

func TestMain(m *testing.M) {
	ownerService = owners.NewOwnerUsace(time.Hour*1, &ownerRepository, &jwtAuth)
	ownerDomain = owners.Domain{
		ID:             1,
		IdentityNumber: "12345678910",
		Name:           "Name Test",
		Password:       "Name1234",
		Email:          "test@example.com",
		BirthOfDate:    "23-10-1996",
		PlaceOfBirth:   "Palopo",
		PhoneNumber:    "12345678",
		Address:        "Jln Bangau",
		District:       "Tamalate",
		City:           "Makassar",
		Province:       "Sulawesi Selatan",
	}
}

func TestLoginOwner(t *testing.T) {
	t.Run("Login Owner", func(t *testing.T) {
		ownerRepository.On("LoginOwner", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(ownerDomain, nil).Once()

		result, err := ownerService.LoginOwner(context.Background(), ownerDomain.Email, ownerDomain.Password)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Get All Invalid", func(t *testing.T) {
		ownerRepository.On("LoginOwner", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(ownerDomain, nil).Once()

		_, err := ownerService.LoginOwner(context.Background(), ownerDomain.Email, ownerDomain.Password)
		assert.NotNil(t, err)
	})
}

func TestGetOwnerByCity(t *testing.T) {
	t.Run("Get Data Owner By City", func(t *testing.T) {
		ownerRepository.On("GetOwnerByCity", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return([]owners.Domain{ownerDomain}, nil).Once()

		result, err := ownerService.GetOwnerByCity(context.Background(), ownerDomain.City)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Get Data Owner By City", func(t *testing.T) {
		ownerRepository.On("GetOwnerByCity", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return([]owners.Domain{ownerDomain}, nil).Once()

		_, err := ownerService.GetOwnerByCity(context.Background(), ownerDomain.City)
		assert.NotNil(t, err)
	})
}

func TestGetAllOwner(t *testing.T) {
	t.Run("Get Data All Owner", func(t *testing.T) {
		ownerRepository.On("GetAllOwner", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return([]owners.Domain{ownerDomain}, nil).Once()

		result, err := ownerService.GetAllOwner(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Get Data Owner By City", func(t *testing.T) {
		ownerRepository.On("GetAllOwner", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return([]owners.Domain{ownerDomain}, nil).Once()

		_, err := ownerService.GetAllOwner(context.Background())
		assert.NotNil(t, err)
	})
}

func TestGetByEmail(t *testing.T) {
	t.Run("Get Data Owner By Email", func(t *testing.T) {
		ownerRepository.On("GetByEmail", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return([]owners.Domain{ownerDomain}, nil).Once()

		result, err := ownerService.GetByEmail(context.Background(), ownerDomain.Email)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Get Data Owner By Email", func(t *testing.T) {
		ownerRepository.On("GetByEmail", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return([]owners.Domain{ownerDomain}, nil).Once()

		_, err := ownerService.GetByEmail(context.Background(), ownerDomain.Email)
		assert.NotNil(t, err)
	})
}

func TestAddOwner(t *testing.T) {
	t.Run("Add New Owner", func(t *testing.T) {
		ownerRepository.On("AddOwner", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return(nil).Once()

		err := ownerService.AddOwner(context.Background(), &ownerDomain)
		assert.Nil(t, err)
	})
}
