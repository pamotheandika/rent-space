package spaces_test

import (
	"RentSpace/businesses/spaces"
	_spacesMock "RentSpace/businesses/spaces/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	spaceRespository _spacesMock.Repository
	spaceService     spaces.Usecase
	spaceDomain      spaces.Domain
)

func TestMain(m *testing.M) {
	spaceService = spaces.NewSpaceUsace(time.Hour*1, &spaceRespository)
	spaceDomain = spaces.Domain{
		ID:          1,
		Name:        "Depan Bank",
		Address:     "Jalan Mangga No.2",
		District:    "Tamalate",
		City:        "Makassar",
		Province:    "Sulawesi Selatan",
		Email:       "jhon@example.com",
		PhoneNumber: "08123456787",
		IDOwner:     1,
		Cost:        350000,
		Status:      1,
	}
}

func TestAddSpace(t *testing.T) {
	t.Run("Add New Space", func(t *testing.T) {
		spaceRespository.On("AddTransaction", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(nil).Once()

		err := spaceService.AddSpace(context.Background(), &spaceDomain)
		assert.Nil(t, err)
	})

	t.Run("Add New Space", func(t *testing.T) {
		spaceRespository.On("AddTransaction", mock.AnythingOfType("string"), mock.AnythingOfType("stirng")).Return(nil).Once()

		err := spaceService.AddSpace(context.Background(), &spaceDomain)
		assert.NotNil(t, err)
	})
}

func TestGetAllSpace(t *testing.T) {
	t.Run("Get All Space", func(t *testing.T) {
		spaceRespository.On("GetAllSpace", mock.AnythingOfType("context.Context"), mock.AnythingOfType("stirng")).Return([]spaces.Domain{spaceDomain}, nil).Once()

		result, err := spaceService.GetAllSpace(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Add New Space", func(t *testing.T) {
		spaceRespository.On("GetAllSpace", mock.AnythingOfType("context.Context"), mock.AnythingOfType("stirng")).Return([]spaces.Domain{spaceDomain}, nil).Once()

		_, err := spaceService.GetAllSpace(context.Background())
		assert.NotNil(t, err)
	})
}
