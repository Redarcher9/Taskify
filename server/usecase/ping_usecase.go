package usecase

import (
	"context"
	"taskify/domain"
	"time"
)

type PingUsecase struct {
	Repository domain.PingRepository
	timeout    int
}

func NewPingUsecase(PingRepository domain.PingRepository, timeout int) domain.PingUsecase {
	return &PingUsecase{
		Repository: PingRepository,
		timeout:    timeout,
	}
}

func (pu *PingUsecase) CheckPing(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, time.Millisecond*100)
	defer cancel()
	return pu.Repository.CheckPing(ctx)
}
