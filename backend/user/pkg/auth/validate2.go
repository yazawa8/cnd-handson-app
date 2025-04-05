package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

func ValidateAuth(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	return token, err
}
