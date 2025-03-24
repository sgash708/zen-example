package handler

import (
	"context"
	"net/http"

	"github.com/unkeyed/unkey/go/pkg/zen"
)

// HelloHandler handles hello world requests
type HelloHandler struct{}

// NewHelloHandler creates a new hello handler
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// RegisterRoutes registers the hello routes with the server
func (h *HelloHandler) RegisterRoutes(server *zen.Server) {
	helloRoute := zen.NewRoute("GET", "/hello", h.Hello)
	server.RegisterRoute(
		[]zen.Middleware{
			zen.WithTracing(),
		},
		helloRoute,
	)
}

// Hello handles hello world requests
func (h *HelloHandler) Hello(ctx context.Context, s *zen.Session) error {
	return s.JSON(http.StatusOK, map[string]string{
		"message": "Hello, world!",
	})
}
