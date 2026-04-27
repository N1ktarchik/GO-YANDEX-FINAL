package http

import "n1ktarchik/go-final/internal/core/domain"

type AuthHandler struct {
	authService AuthService
}

type AuthService interface {
	Login(password string) (string, error)
	GetPath() string
	Compare(hashPassword string) bool
	ValidateJWT(JWT string) (*domain.Claims, error)
}

func NewAuthTransport(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}
