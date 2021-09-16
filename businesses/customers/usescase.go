package customers

import (
	"RentSpace/app/middleware"
	"RentSpace/businesses"
	"RentSpace/helpers/encrypt"
	"context"
	"strings"
	"time"
)

type customerUsecase struct {
	customerRepository Repository
	contextTimeout     time.Duration
	jwtAuth            *middleware.ConfigJwt
}

func NewCustomerUsace(timeout time.Duration, usecase Repository, jwtauth *middleware.ConfigJwt) Usecase {
	return &customerUsecase{
		customerRepository: usecase,
		contextTimeout:     timeout,
		jwtAuth:            jwtauth,
	}
}

func (cu *customerUsecase) GetAllCustomer(ctx context.Context) ([]Domain, error) {
	res, err := cu.customerRepository.GetAllCustomer(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, err
}

func (cu *customerUsecase) LoginCustomer(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	customerDomain, err := cu.customerRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, customerDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := cu.jwtAuth.GenerateToken(customerDomain.IDCostumer, "customer")
	return token, nil
}

func (uc *customerUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	return Domain{}, nil
}

func (cu *customerUsecase) GetByEmail(ctx context.Context, email string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" {
		return Domain{}, businesses.ErrNewsTitleResource
	}

	resp, err := cu.customerRepository.GetByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *customerUsecase) AddCustomer(ctx context.Context, customerDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	
	existedUser, err := uc.customerRepository.GetByEmail(ctx, customerDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return businesses.ErrDuplicateData
	}

	customerDomain.Password, err = encrypt.Hash(customerDomain.Password)
	if err != nil {
		return businesses.ErrInternalServer
	}
	err = uc.customerRepository.AddCustomer(ctx, customerDomain)
	if err != nil {
		return err
	}

	return nil
}
