package generators

import (
	"fmt"
	"strings"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/utils"
)

// GenerateServiceInterface generates a service interface file
func GenerateServiceInterface(name string) error {
	filename := fmt.Sprintf("service/%s_interface.go", strings.ToLower(name))

	// Get module name from go.mod
	moduleName, err := utils.GetModuleName()
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	return utils.WriteFile(filename, templates.ServiceInterfaceTemplate, map[string]interface{}{
		"ProjectName": moduleName,
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

// generateService generates a service layer implementation with business logic operations.
// Handles data transformation between DTOs and models, applies business rules and validation.
// Provides clean interface between handlers and repositories following service pattern.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateService(name string, fields []string) error {
	filename := fmt.Sprintf("service/%s_service.go", strings.ToLower(name))

	// Parse fields for template
	parsedFields := utils.ParseFields(fields)

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
		"Fields":      parsedFields,
	})
}
