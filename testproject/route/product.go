package route

import (
	"<no value>/adapter"
	"<no value>/handler"
	"<no value>/model"
	"<no value>/repository"
	"<no value>/service"

	"github.com/gofiber/fiber/v2"
)

// SetupProductRoutes sets up routes for Product resource
func SetupProductRoutes(api fiber.Router, db *adapter.DatabaseAdapter) {
	// Initialize repository
	productRepo := repository.NewProductRepository(db)

	// Initialize service
	productService := service.NewProductService(productRepo)

	// Initialize handler
	productHandler := handler.NewProductHandler(productService)

	// Setup routes
	productGroup := api.Group("/products")
	productGroup.Get("/", productHandler.FindAll)
	productGroup.Get("/:id", productHandler.FindById)
	productGroup.Post("/", productHandler.Create)
	productGroup.Put("/:id", productHandler.Update)
	productGroup.Delete("/:id", productHandler.Delete)
}
