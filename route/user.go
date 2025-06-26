package route

import (
	"github.com/Oakhouse-Technology/go-to-oakhouse/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes sets up routes for User resource
func SetupUserRoutes(api fiber.Router) {
	// Initialize handler
	userHandler := handler.NewUserHandler()

	// Setup routes
	userGroup := api.Group("/users")
	userGroup.Get("/", userHandler.FindAll)
	userGroup.Get("/:id", userHandler.FindByID)
	userGroup.Post("/", userHandler.Create)
	userGroup.Put("/:id", userHandler.Update)
	userGroup.Delete("/:id", userHandler.Delete)
}
