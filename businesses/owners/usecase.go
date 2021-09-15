package owners

import (
	"RentSpace/businesses"
	"RentSpace/helpers/encrypt"
	"context"
	"strings"
	"time"
)

type ownerUsecase struct {
	ownerRepository Repository
	contextTimeout  time.Duration
}

func NewOwnerUsace(timeout time.Duration, usecase Repository) Usecase {
	return &ownerUsecase{
		ownerRepository: usecase,
		contextTimeout:  timeout,
	}
}

func (cu *ownerUsecase) GetOwnerByCity(ctx context.Context, city string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(city) == "" {
		return Domain{}, businesses.ErrNewsTitleResource
	}

	resp, err := cu.ownerRepository.GetOwnerByCity(ctx, city)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (cu *ownerUsecase) GetAllOwner(ctx context.Context) ([]Domain, error) {
	res, err := cu.ownerRepository.GetAllOwner(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, err
}

func (cu *ownerUsecase) GetByEmail(ctx context.Context, email string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" {
		return Domain{}, businesses.ErrNewsTitleResource
	}

	resp, err := cu.ownerRepository.GetByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

// Create Space Domain
func (uc *ownerUsecase) AddOwner(ctx context.Context, ownerDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.ownerRepository.GetByEmail(ctx, ownerDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}

	if existedUser != (Domain{}) {
		return businesses.ErrDuplicateData
	}

	ownerDomain.Password, err = encrypt.Hash(ownerDomain.Password)
	if err != nil {
		return businesses.ErrInternalServer
	}

	err = uc.ownerRepository.AddOwner(ctx, ownerDomain)
	if err != nil {
		return err
	}

	return nil
}
