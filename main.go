package main

import (
	"context"
	"log"

	"github.com/sgash708/zen-example/application"
	"github.com/sgash708/zen-example/domain/service"
	"github.com/sgash708/zen-example/handler"
	"github.com/sgash708/zen-example/infrastructure/repository"
	"github.com/unkeyed/unkey/go/pkg/otel/logging"
	"github.com/unkeyed/unkey/go/pkg/zen"
	"github.com/unkeyed/unkey/go/pkg/zen/validation"
)

func main() {
	logger := logging.New()

	// Create server
	server, err := zen.New(zen.Config{
		InstanceID: "quickstart-server",
		Logger:     logger,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	// Create validator
	_, err = validation.New()
	if err != nil {
		log.Fatalf("failed to create validator: %v", err)
	}

	// Setup repositories
	userRepo := repository.NewUserRepository()

	// Setup domain services
	userService := service.NewUserService(userRepo)

	// Setup application services
	userApp := application.NewUserApplication(userService)

	// Setup handlers
	helloHandler := handler.NewHelloHandler()
	userHandler := handler.NewUserHandler(userApp)

	// Register routes
	helloHandler.RegisterRoutes(server)
	userHandler.RegisterRoutes(server)

	// Start server
	logger.Info(
		"starting server",
		"address",
		":8080",
	)
	if err := server.Listen(context.Background(), ":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
