package application

import (
	"context"

	"github.com/sgash708/zen-example/domain"
	"github.com/sgash708/zen-example/domain/service"
	"github.com/unkeyed/unkey/go/pkg/fault"
)

// UserApplication is the application layer for user operations
type UserApplication struct {
	userService *service.UserService
}

// NewUserApplication creates a new user application service
func NewUserApplication(userService *service.UserService) *UserApplication {
	return &UserApplication{
		userService: userService,
	}
}

// UserCreateDTO is the data transfer object for user creation
type UserCreateDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponseDTO is the data transfer object for user responses
type UserResponseDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser handles the user creation process
func (a *UserApplication) CreateUser(ctx context.Context, createDTO UserCreateDTO) (*UserResponseDTO, error) {
	user, err := a.userService.CreateUser(ctx, createDTO.Name, createDTO.Email, createDTO.Password)
	if err != nil {
		// Convert domain errors to application errors with fault
		switch err {
		case domain.ErrEmailAlreadyExists:
			return nil, fault.New(
				"email already exists",
				fault.WithTag(fault.BAD_REQUEST),
				fault.WithDesc(
					"a user with this email already exists",
					"そのメールアドレスは使用されています",
				),
			)
		default:
			// Check if it's a validation error from entity creation
			return nil, fault.New(
				"invalid user data",
				fault.WithTag(fault.BAD_REQUEST),
				fault.WithDesc(
					err.Error(),
					"不正なデータです",
				),
			)
		}
	}

	// Map domain entity to response DTO
	return &UserResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
