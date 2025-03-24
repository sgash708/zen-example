package entity

import (
	"errors"
	"regexp"
)

// User represents a user entity in the domain
type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

// NewUser creates a new user with validations
func NewUser(name, email, password string) (*User, error) {
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	// Simple email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	return &User{
		ID:       "", // Will be set by the repository
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
