package generators

import (
	"fmt"
	"strings"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateMiddleware generates HTTP middleware for cross-cutting concerns like authentication, logging, and validation.
// Creates reusable middleware functions that can be applied to routes for consistent request processing.
// Follows standard middleware patterns for easy integration with existing route handlers.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateMiddleware(name string) error {
	filename := fmt.Sprintf("middleware/%s.go", strings.ToLower(name))
	return utils.WriteFile(filename, templates.MiddlewareTemplate, map[string]string{
		"Name": name,
	})
}