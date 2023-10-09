package domain

import "context"

type Loginstruct struct {
	Email    string
	Password string
}

type LoginUsecase interface {
	Login(c context.Context, ls Loginstruct) error
}
