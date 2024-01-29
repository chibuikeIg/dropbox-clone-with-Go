package services

import (
	"time"
	"user-service/internals/core/domain"
	"user-service/internals/core/ports"
)

type UserService struct {
	repo ports.UserDBRepository
}

func NewUserService(repo ports.UserDBRepository) *UserService {
	return &UserService{repo: repo}
}

func (reg *UserService) Store(user *domain.User) (*domain.User, error) {

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := reg.repo.Table("users").Create(user)

	return user, err
}
