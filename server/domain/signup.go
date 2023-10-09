package domain

import "context"

type SignupUsecase interface {
	SignUp(c context.Context, us UserStruct) error
}
