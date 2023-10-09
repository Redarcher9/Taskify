package usecase

import (
	"context"
	"errors"
	"fmt"
	"taskify/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	Repository domain.UserRepository
	timeout    int
}

func NewLoginUsecase(UserRepository domain.UserRepository, timeout int) domain.LoginUsecase {
	return &LoginUsecase{
		Repository: UserRepository,
		timeout:    timeout,
	}
}

func (lu *LoginUsecase) Login(c context.Context, ls domain.Loginstruct) error {
	ctx, cancel := context.WithTimeout(c, time.Millisecond*100)
	defer cancel()
	userexists := lu.Repository.CheckUser(ls.Email)
	if !userexists {
		return errors.New("user does not exist")
	}
	fmt.Println(ctx.Done())
	password, err := lu.Repository.GetPassword(ls.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(ls.Password))
	if err != nil {
		return errors.New("incorrect credentials")
	}
	return err
}
