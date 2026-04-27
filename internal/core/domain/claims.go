package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	HashPass string
	jwt.RegisteredClaims
}
