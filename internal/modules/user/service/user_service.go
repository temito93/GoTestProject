package service

import (
	"context"
	"goproject/internal/modules/user/entity"
	"goproject/internal/modules/user/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (*entity.UserEntity, error) {
	// Create a new UserEntity instance
	user := &entity.UserEntity{Name: name, Email: email}

	// Call the repository to create the user
	userData, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// Return the user data with the generated ID
	return userData, nil
}
