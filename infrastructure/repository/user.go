package repository

import (
	"context"
	"strconv"
	"sync"

	"github.com/sgash708/zen-example/domain"
	"github.com/sgash708/zen-example/domain/entity"
)

// UserRepository implements UserRepository with in-memory storage
type UserRepository struct {
	users  map[string]*entity.User
	emails map[string]string // email -> id mapping
	nextID int
	mu     sync.RWMutex
}

// NewUserRepository creates a new in-memory user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[string]*entity.User),
		emails: make(map[string]string),
		nextID: 1,
	}
}

// Create stores a new user in memory
func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if email exists
	if _, exists := r.emails[user.Email]; exists {
		return nil, domain.ErrEmailAlreadyExists
	}

	// Generate ID
	id := generateID(r.nextID)
	r.nextID++

	// Clone and store user
	newUser := &entity.User{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	r.users[id] = newUser
	r.emails[user.Email] = id

	return newUser, nil
}

// FindByID retrieves a user by ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	// Return a copy to prevent accidental modification
	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// FindByEmail retrieves a user by email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, exists := r.emails[email]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	user := r.users[id]
	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// Helper to generate an ID string
func generateID(id int) string {
	return "user_" + strconv.Itoa(id)
}
