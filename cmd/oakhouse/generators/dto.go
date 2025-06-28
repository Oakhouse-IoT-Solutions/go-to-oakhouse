package generators

import (
	"fmt"
	"os"
	"strings"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateDTO generates Data Transfer Objects for Create, Update, and Get operations.
// Creates separate DTOs with proper validation tags for request/response data transformation.
// Ensures clean separation between API contracts and internal data models.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateDTO(name string, fields []string) error {
	dtoDir := fmt.Sprintf("dto/%s", strings.ToLower(name))
	if err := os.MkdirAll(dtoDir, 0755); err != nil {
		return fmt.Errorf("failed to create DTO directory: %w", err)
	}

	// Parse fields for template
	parsedFields := utils.ParseFields(fields)
	
	// Get module name from go.mod
	moduleName, err := utils.GetModuleName()
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	dtoTemplates := map[string]string{
		"create": templates.CreateDtoTemplate,
		"update": templates.UpdateDtoTemplate,
		"get":    templates.GetDtoTemplate,
	}

	for dtoType, tmpl := range dtoTemplates {
		filename := fmt.Sprintf("%s/%s_%s_dto.go", dtoDir, dtoType, strings.ToLower(name))
		if err := utils.WriteFile(filename, tmpl, map[string]interface{}{
			"ProjectName": moduleName,
			"ModelName":   name,
			"PackageName": strings.ToLower(name),
			"VarName":     strings.ToLower(name),
			"Type":        strings.Title(dtoType),
			"Fields":      parsedFields,
		}); err != nil {
			return err
		}
	}

	return nil
}