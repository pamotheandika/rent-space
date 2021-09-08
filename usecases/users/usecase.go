package users

import "time"

type userUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
}

func NewUserUsace(usecase Repository, timeout time.Duration) Usecase {
	return &userUsecase{
		userRepository: usecase,
		contextTimeout: timeout,
	}
}

// Create user Domain
func (uuc *userUsecase) CreateUser(user Domain) (int, error) {
	// TODO bussiness logic
	result, err := uuc.userRepository.UpsertUser(user)
	if err != nil {
		return 0, err
	}
	return result.ID, nil
}

// Add Device when user register the device
func (uuc *userUsecase) AddDevice(device Device, userID int) error {
	return nil
}
