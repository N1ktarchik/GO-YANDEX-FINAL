package service

import (
	"n1ktarchik/go-final/internal/core/domain"
	core_errors "n1ktarchik/go-final/internal/core/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *AuthService) CreateJWT(HashPassword string) (string, error) {
	secret := s.secretKey

	claims := domain.Claims{
		HashPass: HashPassword,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(8 * time.Hour))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "todo-App",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", core_errors.ServerError("creating JWT faild")
	}

	return token, nil
}

func (s *AuthService) ValidateJWT(JWT string) (*domain.Claims, error) {
	claims := &domain.Claims{}

	token, err := jwt.ParseWithClaims(JWT, claims, func(t *jwt.Token) (any, error) {

		if t.Method != jwt.SigningMethodHS256 {
			return nil, core_errors.InvalidJWT()
		}

		return s.secretKey, nil
	})

	if err != nil {
		return nil, core_errors.ServerError("jwt validation failed")
	}

	if token.Valid {

		return claims, nil

	}

	return nil, core_errors.InvalidJWT()
}
