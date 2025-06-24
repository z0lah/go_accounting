package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type TokenParser interface {
	Parse(tokenStr string) (*jwt.Token, jwt.MapClaims, error)
}

func (j *jwtTokenGenerator) Parse(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}
	return nil, nil, errors.New("invalid token")
}
