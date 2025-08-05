package generators

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateScope generates GORM scopes for advanced query filtering and reusable query logic.
// Creates scope functions for common filtering patterns like field-specific filters and date ranges.
// Promotes code reuse and maintains clean separation of query logic.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
// GenerateFieldFilter generates a specific field filter function and appends it to the filter.go file
func GenerateFieldFilter(modelName, fieldName, fieldType string) error {
	scopeDir := fmt.Sprintf("scope/%s", strings.ToLower(modelName))
	filterFilename := fmt.Sprintf("%s/filter.go", scopeDir)

	// Check if filter.go exists
	if _, err := os.Stat(filterFilename); os.IsNotExist(err) {
		// Create the base filter.go first
		if err := GenerateScope(modelName, ""); err != nil {
			return err
		}
	}

	// Read current content
	currentContent, err := os.ReadFile(filterFilename)
	if err != nil {
		return err
	}

	// Check if this filter function already exists
	if strings.Contains(string(currentContent), fmt.Sprintf("FilterBy%s", fieldName)) {
		return nil // Filter already exists
	}

	// Generate filter function using template
	templateData := map[string]interface{}{
		"PackageName": strings.ToLower(modelName),
		"ModelName":   modelName,
		"FieldName":   fieldName,
		"ParamName":   strings.ToLower(fieldName),
		"ParamType":   fieldType,
		"ColumnName":  strings.ToLower(fieldName),
	}

	filterFunc, err := renderTemplate(templates.ScopeTemplate, templateData)
	if err != nil {
		return fmt.Errorf("failed to render scope template: %w", err)
	}

	// Append the new filter function
	newContent := string(currentContent) + "\n" + filterFunc
	return os.WriteFile(filterFilename, []byte(newContent), 0644)
}

// GenerateBaseScope creates the base_scope.go file with common utility functions
func GenerateBaseScope() error {
	baseScopeFilename := "scope/base_scope.go"

	// Check if base_scope.go already exists
	if _, err := os.Stat(baseScopeFilename); os.IsNotExist(err) {
		// Create scope directory if it doesn't exist
		if err := os.MkdirAll("scope", 0755); err != nil {
			return fmt.Errorf("failed to create scope directory: %w", err)
		}

		// Generate base scope content with consistent pointer types
		baseScopeContent := `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package scope

import (
	"time"

	"gorm.io/gorm"
)

// DateRangeScope applies date range filtering with time.Time values for consistency
func DateRangeScope(startDate, endDate time.Time, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !startDate.IsZero() {
			db = db.Where(column+" >= ?", startDate)
		}
		if !endDate.IsZero() {
			db = db.Where(column+" <= ?", endDate)
		}
		return db
	}
}

// SearchScope applies text search filtering
func SearchScope(search, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			return db.Where(column+" ILIKE ?", "%"+search+"%")
		}
		return db
	}
}

// StatusScope applies status filtering
func StatusScope(status, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where(column+" = ?", status)
		}
		return db
	}
}
`

		if err := os.WriteFile(baseScopeFilename, []byte(baseScopeContent), 0644); err != nil {
			return err
		}
	}

	return nil
}

func GenerateScope(modelName, scopeName string) error {
	scopeDir := fmt.Sprintf("scope/%s", strings.ToLower(modelName))
	if err := os.MkdirAll(scopeDir, 0755); err != nil {
		return fmt.Errorf("failed to create scope directory: %w", err)
	}

	// Generate the filter.go file with field-specific filter functions
	filterFilename := fmt.Sprintf("%s/filter.go", scopeDir)
	_, err := utils.GetModuleName()
	if err != nil {
		// fallback handled in template
	}

	// Check if filter.go already exists
	if _, err := os.Stat(filterFilename); os.IsNotExist(err) {
		// Generate filter content using template
		templateData := map[string]interface{}{
			"PackageName": strings.ToLower(modelName),
			"ModelName":   modelName,
			"FieldName":   "DateRange",
			"ColumnName":  "created_at",
		}

		filterContent, err := renderTemplate(templates.DateRangeFilterTemplate, templateData)
		if err != nil {
			return fmt.Errorf("failed to render date range filter template: %w", err)
		}

		if err := os.WriteFile(filterFilename, []byte(filterContent), 0644); err != nil {
			return err
		}
	}

	return nil
}

// GeneratePaginationScope generates pagination scope functions and appends them to the filter.go file
func GeneratePaginationScope(modelName string) error {
	scopeDir := fmt.Sprintf("scope/%s", strings.ToLower(modelName))
	filterFilename := fmt.Sprintf("%s/filter.go", scopeDir)

	// Check if filter.go exists
	if _, err := os.Stat(filterFilename); os.IsNotExist(err) {
		// Create the base filter.go first
		if err := GenerateScope(modelName, ""); err != nil {
			return err
		}
	}

	// Read current content
	currentContent, err := os.ReadFile(filterFilename)
	if err != nil {
		return err
	}

	// Check if pagination functions already exist
	if strings.Contains(string(currentContent), "PaginationScope") {
		return nil // Pagination already exists
	}

	// Generate pagination functions using template
	templateData := map[string]interface{}{
		"PackageName": strings.ToLower(modelName),
		"ModelName":   modelName,
	}

	paginationFunc, err := renderTemplate(templates.PaginationScopeFunctionOnlyTemplate, templateData)
	if err != nil {
		return fmt.Errorf("failed to render pagination scope template: %w", err)
	}

	// Append the pagination functions
	newContent := string(currentContent) + "\n" + paginationFunc
	return os.WriteFile(filterFilename, []byte(newContent), 0644)
}

// GenerateAdvancedDateRangeFilter generates advanced date range filter using template
func GenerateAdvancedDateRangeFilter(modelName string) error {
	scopeDir := fmt.Sprintf("scope/%s", strings.ToLower(modelName))
	filterFilename := fmt.Sprintf("%s/filter.go", scopeDir)

	// Check if filter.go exists
	if _, err := os.Stat(filterFilename); os.IsNotExist(err) {
		// Create the base filter.go first
		if err := GenerateScope(modelName, ""); err != nil {
			return err
		}
	}

	// Read current content
	currentContent, err := os.ReadFile(filterFilename)
	if err != nil {
		return err
	}

	// Check if advanced date range filter already exists
	if strings.Contains(string(currentContent), "DateRangeFilter") {
		return nil // Advanced date range filter already exists
	}

	// Generate advanced date range filter using template
	templateData := map[string]interface{}{
		"PackageName": strings.ToLower(modelName),
		"ModelName":   modelName,
	}

	advancedFilterFunc, err := renderTemplate(templates.AdvancedDateRangeFilterFunctionOnlyTemplate, templateData)
	if err != nil {
		return fmt.Errorf("failed to render advanced date range filter template: %w", err)
	}

	// Append the advanced date range filter
	newContent := string(currentContent) + "\n" + advancedFilterFunc
	return os.WriteFile(filterFilename, []byte(newContent), 0644)
}

// GenerateCompleteScope generates a complete scope package with all common filters
func GenerateCompleteScope(modelName string, fields []string) error {
	// Generate base scope first
	if err := GenerateScope(modelName, ""); err != nil {
		return fmt.Errorf("failed to generate base scope: %w", err)
	}

	// Generate pagination scope
	if err := GeneratePaginationScope(modelName); err != nil {
		return fmt.Errorf("failed to generate pagination scope: %w", err)
	}

	// Generate field-specific filters
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			continue // Skip invalid field format
		}
		fieldName := parts[0]
		fieldType := parts[1]

		if err := GenerateFieldFilter(modelName, fieldName, fieldType); err != nil {
			return fmt.Errorf("failed to generate field filter for %s: %w", fieldName, err)
		}
	}

	return nil
}

// GenerateScopeWithOptions generates scope with specific options
type ScopeOptions struct {
	IncludePagination        bool
	IncludeAdvancedDateRange bool
	Fields                   []string // Format: "fieldName:fieldType"
}

func GenerateScopeWithOptions(modelName string, options ScopeOptions) error {
	// Generate base scope first
	if err := GenerateScope(modelName, ""); err != nil {
		return fmt.Errorf("failed to generate base scope: %w", err)
	}

	// Generate pagination scope if requested
	if options.IncludePagination {
		if err := GeneratePaginationScope(modelName); err != nil {
			return fmt.Errorf("failed to generate pagination scope: %w", err)
		}
	}

	// Generate advanced date range filter if requested
	if options.IncludeAdvancedDateRange {
		if err := GenerateAdvancedDateRangeFilter(modelName); err != nil {
			return fmt.Errorf("failed to generate advanced date range filter: %w", err)
		}
	}

	// Generate field-specific filters
	for _, field := range options.Fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			continue // Skip invalid field format
		}
		fieldName := parts[0]
		fieldType := parts[1]

		if err := GenerateFieldFilter(modelName, fieldName, fieldType); err != nil {
			return fmt.Errorf("failed to generate field filter for %s: %w", fieldName, err)
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
