package service

type AuthService struct {
	secretKey []byte
	password  string
}

func NewAuthService(secret, password string) *AuthService {
	service := &AuthService{
		secretKey: []byte(secret),
	}

	service.password = service.Hash(password)

	return service

}
