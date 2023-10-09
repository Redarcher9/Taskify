package usecase

import (
	"context"
	"errors"
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
	userexists := su.Repository.CheckUser(us.Email)
	if userexists {
		return errors.New("user already exists")
	}
	fmt.Println(ctx.Done())
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(us.Password), 8)
	if err != nil {
		return err
	}
	us.Password = string(passwordBytes)
	err = su.Repository.CreateUser(us)
	if err != nil {
		return err
	}
	return err
}
