package owners

import (
	"RentSpace/app/middleware"
	"RentSpace/businesses"
	"RentSpace/helpers/encrypt"
	"context"
	"strings"
	"time"
)

type ownerUsecase struct {
	ownerRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJwt
}

func NewOwnerUsace(timeout time.Duration, usecase Repository, jwtauth *middleware.ConfigJwt) Usecase {
	return &ownerUsecase{
		ownerRepository: usecase,
		contextTimeout:  timeout,
		jwtAuth:         jwtauth,
	}
}

func (oc *ownerUsecase) LoginOwner(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, oc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	resp, err := oc.ownerRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, resp.Password) {
		return "", businesses.ErrInternalServer
	}

	token := oc.jwtAuth.GenerateToken(resp.ID, "owner")
	return token, nil

}

func (cu *ownerUsecase) GetOwnerByCity(ctx context.Context, city string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(city) == "" {
		return []Domain{}, businesses.ErrNewsTitleResource
	}

	resp, err := cu.ownerRepository.GetOwnerByCity(ctx, city)
	if err != nil {
		return []Domain{}, err
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
