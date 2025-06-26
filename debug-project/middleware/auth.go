package middleware

import (
	"strings"

	"debug-project/config"
	"debug-project/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := extractToken(c)
		if token == "" {
			return util.SendError(c, fiber.StatusUnauthorized, "Missing or invalid token", nil)
		}

		claims, err := validateToken(token)
		if err != nil {
			return util.SendError(c, fiber.StatusUnauthorized, "Invalid token", err.Error())
		}

		// Store user info in context
		c.Locals("user_id", claims["user_id"])
		c.Locals("user_email", claims["email"])

		return c.Next()
	}
}

func extractToken(c *fiber.Ctx) string {
	auth := c.Get("Authorization")
	if auth == "" {
		return ""
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	cfg := config.Load()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

func RateLimit(cfg *config.Config) fiber.Handler {
	// Implementation would use a rate limiting library
	// For now, just return a pass-through middleware
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
