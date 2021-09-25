package cities

import (
	"context"
	"time"
)

type citiesUsecase struct {
	citiesRepository Repository
	contextTimeout   time.Duration
}

func NewCitiesUsecase(timeout time.Duration, usecase Repository) Usecase {
	return &citiesUsecase{
		citiesRepository: usecase,
		contextTimeout: timeout,
	}
}

func(cu *citiesUsecase) GetCities(ctx context.Context)  ([]Domain, error) {
	res, err := cu.citiesRepository.GetCities(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, err
}


