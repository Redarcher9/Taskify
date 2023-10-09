package domain

import "context"

type Loginstruct struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type LoginUsecase interface {
	Login(c context.Context, ls Loginstruct) error
}
