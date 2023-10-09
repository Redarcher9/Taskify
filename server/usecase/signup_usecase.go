package usecase

import (
	"context"
	"fmt"
	"taskify/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type SignupUsecase struct {
	Repository domain.UserRepository
	timeout    int
}

func NewSignupUsecase(UserRepository domain.UserRepository, timeout int) domain.SignupUsecase {
	return &SignupUsecase{
		Repository: UserRepository,
		timeout:    timeout,
	}
}

func (su *SignupUsecase) SignUp(c context.Context, us domain.UserStruct) error {
	ctx, cancel := context.WithTimeout(c, time.Millisecond*100)
	defer cancel()
	err := su.Repository.CheckUser(us.Email)
	if err != nil {
		return err
	}
	fmt.Println(ctx.Done())
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(us.Password), 8)
	us.Password = string(passwordBytes)
	err = su.Repository.CreateUser(us)
	if err != nil {
		return err
	}
	return err
}
