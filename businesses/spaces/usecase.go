package spaces

import (
	"context"
	"time"
)

type spaceUsecase struct {
	spaceRepository Repository
	contextTimeout  time.Duration
}

func NewSpaceUsace(timeout time.Duration, usecase Repository) Usecase {
	return &spaceUsecase{
		spaceRepository: usecase,
		contextTimeout:  timeout,
	}
}

func (sc *spaceUsecase) UpdateStatusSpace(ctx context.Context, IDSpace int) error {
	err := sc.spaceRepository.UpdateStatusSpace(ctx, IDSpace)

	if err != nil {
		return err
	}

	return err
}

func (sc *spaceUsecase) GetAllSpace(ctx context.Context) ([]Domain, error) {
	res, err := sc.spaceRepository.GetAllSpace(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return res, err
}

func (sc *spaceUsecase) AddSpace(ctx context.Context, spaceDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, sc.contextTimeout)
	defer cancel()

	err := sc.spaceRepository.AddSpace(ctx, spaceDomain)
	if err != nil {
		return err
	}

	return nil
}
