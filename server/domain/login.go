package domain

import (
	"context"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Loginstruct struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type MyCustomClaims struct {
	Name string
	jwt.RegisteredClaims
}

type LoginUsecase interface {
	Login(c context.Context, ls Loginstruct) error
}
