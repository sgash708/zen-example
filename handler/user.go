package handler

import (
	"context"
	"net/http"

	"github.com/sgash708/zen-example/application"
	"github.com/unkeyed/unkey/go/pkg/fault"
	"github.com/unkeyed/unkey/go/pkg/zen"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userApp *application.UserApplication
}

// NewUserHandler creates a new user handler
func NewUserHandler(userApp *application.UserApplication) *UserHandler {
	return &UserHandler{
		userApp: userApp,
	}
}

// RegisterRoutes registers the user routes with the server
func (h *UserHandler) RegisterRoutes(server *zen.Server) {
	createUserRoute := zen.NewRoute("POST", "/users", h.CreateUser)
	server.RegisterRoute(
		[]zen.Middleware{
			zen.WithTracing(),
		},
		createUserRoute,
	)
}

// CreateUser handles user creation requests
func (h *UserHandler) CreateUser(ctx context.Context, s *zen.Session) error {
	var req application.UserCreateDTO
	if err := s.BindBody(&req); err != nil {
		return fault.New("invalid request body",
			fault.WithTag(fault.BAD_REQUEST),
			fault.WithDesc(
				"could not parse request body",
				"リクエストボディを解析できませんでした"),
		)
	}

	// Process request through application layer
	user, err := h.userApp.CreateUser(ctx, req)
	if err != nil {
		return fault.Wrap(err)
	}

	return s.JSON(http.StatusCreated, user)
}
