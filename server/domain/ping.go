package domain

import "context"

type PingRepository interface {
	CheckPing(c context.Context) error
}

type PingUsecase interface {
	CheckPing(c context.Context) error
}
