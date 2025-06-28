package generators

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateScope generates GORM scopes for advanced query filtering and reusable query logic.
// Creates scope functions for common filtering patterns like date ranges, status filters, and pagination.
// Promotes code reuse and maintains clean separation of query logic.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateScope(modelName, scopeName string) error {
	scopeDir := fmt.Sprintf("scope/%s", strings.ToLower(modelName))
	if err := os.MkdirAll(scopeDir, 0755); err != nil {
		return fmt.Errorf("failed to create scope directory: %w", err)
	}

	// Generate the specific scope file
	filename := fmt.Sprintf("%s/%s.go", scopeDir, strings.ToLower(scopeName))
	if err := utils.WriteFile(filename, templates.ScopeTemplate, map[string]interface{}{
		"ModelName":   modelName,
		"PackageName": strings.ToLower(modelName),
		"ScopeName":   scopeName,
		"FieldName":   scopeName,
	}); err != nil {
		return err
	}

	// Generate the filter.go file with date range and pagination scopes
	filterFilename := fmt.Sprintf("%s/filter.go", scopeDir)
	projectName, err := utils.GetModuleName()
	if err != nil {
		projectName = "your-project" // fallback
	}

	// Check if filter.go already exists
	if _, err := os.Stat(filterFilename); os.IsNotExist(err) {
		if err := utils.WriteFile(filterFilename, templates.DateRangeFilterTemplate, map[string]interface{}{
			"ModelName":   modelName,
			"PackageName": strings.ToLower(modelName),
			"ProjectName": projectName,
		}); err != nil {
			return err
		}

		// Append pagination scope to the filter.go file
		paginationContent, err := renderTemplate(templates.PaginationScopeTemplate, map[string]interface{}{
			"ModelName":   modelName,
			"PackageName": strings.ToLower(modelName),
		})
		if err != nil {
			return err
		}

		// Read the current filter.go content
		currentContent, err := os.ReadFile(filterFilename)
		if err != nil {
			return err
		}

		// Append pagination scope
		newContent := string(currentContent) + "\n\n" + paginationContent
		if err := os.WriteFile(filterFilename, []byte(newContent), 0644); err != nil {
			return err
		}
	}

	return nil
}

// renderTemplate is a helper function to render templates with data
func renderTemplate(tmpl string, data interface{}) (string, error) {
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}