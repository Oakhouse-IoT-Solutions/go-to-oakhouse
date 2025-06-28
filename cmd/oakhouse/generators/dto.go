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
func GenerateDTO(name string) error {
	dtoDir := fmt.Sprintf("dto/%s", strings.ToLower(name))
	if err := os.MkdirAll(dtoDir, 0755); err != nil {
		return fmt.Errorf("failed to create DTO directory: %w", err)
	}

	dtoTemplates := map[string]string{
		"create": templates.CreateDtoTemplate,
		"update": templates.UpdateDtoTemplate,
		"get":    templates.GetDtoTemplate,
	}

	for dtoType, tmpl := range dtoTemplates {
		filename := fmt.Sprintf("%s/%s_%s_dto.go", dtoDir, dtoType, strings.ToLower(name))
		if err := utils.WriteFile(filename, tmpl, map[string]interface{}{
			"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
			"ModelName":   name,
			"PackageName": strings.ToLower(name),
			"VarName":     strings.ToLower(name),
			"Type":        strings.Title(dtoType),
		}); err != nil {
			return err
		}
	}

	return nil
}