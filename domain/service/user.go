package service

import (
	"context"

	"github.com/sgash708/zen-example/domain"
	"github.com/sgash708/zen-example/domain/entity"
	"github.com/sgash708/zen-example/domain/repository"
)

// UserService contains domain logic for users
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service with the given repository
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user with the given details
func (s *UserService) CreateUser(ctx context.Context, name, email, password string) (*entity.User, error) {
	// Check if user with email already exists
	existingUser, _ := s.userRepo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	// Create and validate user
	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	// Save user
	return s.userRepo.Create(ctx, user)
}