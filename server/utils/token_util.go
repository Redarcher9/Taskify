package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

/*
Description: Creates JWT token for authentication

Params: Gin context, email string, HMAC secret

Returns: signedtoken and error
*/

func CreateToken(c *gin.Context, email string, secret string) (string, error) {

	type MyCustomClaims struct {
		Name string
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		email,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return ss, err
}

/*
Description: Validates JWT token for authentication

Params: Gin context, email string, HMAC secret

Returns: nil-if valid, error-nil if valid or error
*/

func ValidateToken(c *gin.Context, tokenString string, secret string) error {
	type MyCustomClaims struct {
		Name string
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return nil
	} else {
		return err
	}
}
