package generators

import (
	"fmt"
	"strings"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateRepository generates a repository implementation with full CRUD operations.
// Includes context support, GORM scopes, pagination, and proper error handling.
// Follows repository pattern for clean separation between business logic and data access.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateRepository(name string) error {
	filename := fmt.Sprintf("repository/%s_repo.go", strings.ToLower(name))
	// Get module name from go.mod
	moduleName, err := utils.GetModuleName()
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	return utils.WriteFile(filename, templates.RepositoryInterfaceTemplate+"\n\n"+templates.RepositoryImplTemplate, map[string]interface{}{
		"ProjectName": moduleName,
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}
