package services

import (
	"api-gateway/internals/core/domain"
	"api-gateway/internals/core/ports"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	repo ports.ApiGatewayRepository
}

func NewRegistrationService(repo ports.ApiGatewayRepository) *RegistrationService {

	return &RegistrationService{
		repo: repo,
	}
}

func (reg *RegistrationService) Store(user *domain.User) (*domain.User, error) {

	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// hash password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(password)

	if err != nil {
		return nil, err
	}

	_, err = reg.repo.Table("users").Create(user)

	return user, err
}

func (reg *RegistrationService) CheckUserExists(email string) bool {

	var users []domain.User

	reg.repo.Table("users").Find([]string{
		"email",
		email,
	}, &users)

	return len(users) != 0
}
