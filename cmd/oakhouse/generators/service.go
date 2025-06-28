package generators

import (
	"fmt"
	"strings"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateService generates a service layer implementation with business logic operations.
// Handles data transformation between DTOs and models, applies business rules and validation.
// Provides clean interface between handlers and repositories following service pattern.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateService(name string) error {
	filename := fmt.Sprintf("service/%s_service.go", strings.ToLower(name))
	return utils.WriteFile(filename, templates.ServiceImplTemplate, map[string]interface{}{
		"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}