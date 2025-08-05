package generators

import (
	"fmt"
	"strings"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateHandler generates a REST API handler with full CRUD endpoints.
// Creates HTTP handlers for Create, Read, Update, Delete operations with proper status codes,
// request validation, error handling, and JSON responses following REST conventions.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateHandler(name string) error {
	filename := fmt.Sprintf("handler/%s_handler.go", strings.ToLower(name))
	// Get module name from go.mod
	moduleName, err := utils.GetModuleName()
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	return utils.WriteFile(filename, templates.HandlerTemplate, map[string]interface{}{
		"ProjectName": moduleName,
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

func GenerateSimpleHandler(name string) error {
	filename := fmt.Sprintf("handler/%s_handler.go", strings.ToLower(name))
	// Get module name from go.mod
	moduleName, err := utils.GetModuleName()
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	return utils.WriteFile(filename, templates.SimpleHandlerTemplate, map[string]interface{}{
		"ProjectName": moduleName,
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}
