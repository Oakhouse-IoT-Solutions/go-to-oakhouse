// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware handles authentication
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip auth for health check and public routes
		if c.Path() == "/health" || c.Path() == "/" {
			return c.Next()
		}

		// Add your authentication logic here
		// For now, just pass through
		return c.Next()
	}
}
