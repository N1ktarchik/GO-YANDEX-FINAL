package service

import core_errors "n1ktarchik/go-final/internal/core/errors"

func (s *AuthService) Login(password string) (string, error) {
	hashPass := s.Hash(password)

	if !s.Compare(hashPass) {
		return "", core_errors.InvalidPassword()
	}

	token, err := s.CreateJWT(hashPass)
	if err != nil {
		return "", err
	}

	return token, nil
}
