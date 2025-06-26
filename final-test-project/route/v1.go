package route

import (
	"final-test-project/adapter"
	"final-test-project/handler"
	"final-test-project/repository"
	"final-test-project/service"
	"final-test-project/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupV1Routes(api fiber.Router, db *adapter.DatabaseAdapter) {
	v1 := api.Group("/v1")

	// Initialize repositories
	// userRepo := repository.NewUserRepository(db)

	// Initialize services
	// userService := service.NewUserService(userRepo)

	// Initialize handlers
	// userHandler := handler.NewUserHandler(userService)

	// Public routes
	public := v1.Group("/")
	// public.Post("/auth/login", authHandler.Login)
	// public.Post("/auth/register", authHandler.Register)

	// Protected routes
	protected := v1.Group("/", middleware.AuthRequired())
	// protected.Get("/users", userHandler.FindAll)
	// protected.Get("/users/:id", userHandler.FindById)
	// protected.Post("/users", userHandler.Create)
	// protected.Put("/users/:id", userHandler.Update)
	// protected.Delete("/users/:id", userHandler.Delete)
}
