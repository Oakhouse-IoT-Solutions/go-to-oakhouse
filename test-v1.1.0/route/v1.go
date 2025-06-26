package route

import (
	"test-v1.1.0/adapter"
	// "test-v1.1.0/handler"
	// "test-v1.1.0/repository"
	// "test-v1.1.0/service"
	// "test-v1.1.0/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupV1Routes(api fiber.Router, db *adapter.DatabaseAdapter) {
	v1 := api.Group("/v1")

	// Health check endpoint
	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Server is running",
		})
	})

	// Initialize repositories (uncomment when needed)
	// userRepo := repository.NewUserRepository(db)

	// Initialize services (uncomment when needed)
	// userService := service.NewUserService(userRepo)

	// Initialize handlers (uncomment when needed)
	// userHandler := handler.NewUserHandler(userService)

	// Public routes (uncomment when needed)
	// public := v1.Group("/")
	// public.Post("/auth/login", authHandler.Login)
	// public.Post("/auth/register", authHandler.Register)

	// Protected routes (uncomment when needed)
	// protected := v1.Group("/", middleware.AuthRequired())
	// protected.Get("/users", userHandler.FindAll)
	// protected.Get("/users/:id", userHandler.FindById)
	// protected.Post("/users", userHandler.Create)
	// protected.Put("/users/:id", userHandler.Update)
	// protected.Delete("/users/:id", userHandler.Delete)
}
