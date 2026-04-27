package service

import (
	"crypto/sha256"
	"encoding/hex"
)

func (s *AuthService) Hash(password string) string {

	hashPassword := sha256.Sum256([]byte(password))

	return hex.EncodeToString(hashPassword[:])
}

func (s *AuthService) Compare(hashPassword string) bool {
	return hashPassword == s.password
}

func (s *AuthService) GetPath() string {
	return s.password
}
