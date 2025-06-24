package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenGenerator interface {
	Generate(userID string, email string, role string) (string, error)
}

type jwtTokenGenerator struct {
	secretKey string
	expiry    time.Duration
}
type TokenService interface {
	TokenGenerator
	TokenParser
}

func NewJWTGenerator(secretKey string, expiry time.Duration) TokenService {
	return &jwtTokenGenerator{
		secretKey: secretKey,
		expiry:    expiry,
	}
}

func (j *jwtTokenGenerator) Generate(userID, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(j.expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}
