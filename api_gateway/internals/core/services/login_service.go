package services

import (
	"api-gateway/internals/core/domain"
	"api-gateway/internals/core/ports"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo ports.ApiGatewayRepository
}

func NewLoginService(repo ports.ApiGatewayRepository) *LoginService {

	return &LoginService{repo: repo}
}

func (ls *LoginService) Authenticate(email string, password string) (*domain.User, error) {

	var users []domain.User

	ls.repo.Table("users").Find([]string{
		"email",
		email,
	}, &users)

	if len(users) == 0 {
		return nil, errors.New("the credential provided doesn't match our records")
	}

	user := users[0]

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("the credential provided doesn't match our records")
	}

	return &user, nil
}
