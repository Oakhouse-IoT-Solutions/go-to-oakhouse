// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// SuccessResponse returns a standardized success response
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(fiber.Map{
		"requestId": uuid.New(),
		"success":   true,
		"message":   message,
		"data":      data,
	})
}

// ErrorResponse returns a standardized error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err error) error {
	response := fiber.Map{
		"requestId": uuid.New(),
		"success":   false,
		"message":   message,
	}

	if err != nil {
		response["error"] = err.Error()
	}

	return c.Status(statusCode).JSON(response)
}

// PaginatedResponse returns a standardized paginated response
func PaginatedResponse(c *fiber.Ctx, message string, data interface{}, pagination interface{}) error {
	return c.JSON(fiber.Map{
		"requestId":  uuid.New(),
		"success":    true,
		"message":    message,
		"data":       data,
		"pagination": pagination,
	})
}
