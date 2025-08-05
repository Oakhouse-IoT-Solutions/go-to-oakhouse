package generators

import (
	"fmt"
	"strings"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateModel generates a GORM model with UUID primary key, timestamps, and soft delete support.
// Creates a struct with proper GORM tags and JSON serialization for database operations.
// Fields are parsed and mapped to appropriate Go types with validation tags.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateModel(name string, fields []string) error {
	filename := fmt.Sprintf("model/%s.go", strings.ToLower(name))
	return utils.WriteFile(filename, templates.ModelTemplate, map[string]interface{}{
		"ModelName": name,
		"TableName": strings.ToLower(name) + "s",
		"Fields":    utils.ParseFields(fields),
	})
}
