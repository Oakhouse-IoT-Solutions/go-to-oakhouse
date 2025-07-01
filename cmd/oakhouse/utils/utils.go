// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// Version represents the current version of Go To Oakhouse
const Version = "1.29.0"

// Field represents a struct field with metadata for code generation
type Field struct {
	Name      string
	Type      string
	Tag       string
	GormTag   string
	JsonTag   string
	QueryType string
	QueryTag  string
}

// WriteFile creates files from Go templates with dynamic data injection.
// Handles directory creation, template parsing, and file generation with proper error handling.
// Core utility function used by all code generators for consistent file creation.
func WriteFile(filename, tmpl string, data interface{}) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Parse template
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	// Execute template
	if err := t.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

// ParseFields parses field definitions from command line arguments into structured Field objects.
// Converts string field definitions (name:type format) into Field structs with proper Go types,
// GORM tags, and JSON tags for database and API serialization.
func ParseFields(fields []string) []Field {
	var result []Field
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) == 2 {
			fieldName := strings.Title(parts[0])
			fieldType := MapGoType(parts[1])
			lowerName := strings.ToLower(parts[0])

			result = append(result, Field{
				Name:      fieldName,
				Type:      fieldType,
				Tag:       fmt.Sprintf(`json:"%s" gorm:"column:%s"`, lowerName, lowerName),
				GormTag:   fmt.Sprintf("column:%s", lowerName),
				JsonTag:   lowerName,
				QueryType: "*" + fieldType,
				QueryTag:  lowerName,
			})
		}
	}
	return result
}

// MapGoType maps string type names to their corresponding Go type declarations.
// Provides type mapping from user-friendly names to proper Go types,
// supporting common types like string, int, bool, time, uuid with sensible defaults.
func MapGoType(t string) string {
	switch strings.ToLower(t) {
	case "string":
		return "string"
	case "int":
		return "int"
	case "int64":
		return "int64"
	case "float64":
		return "float64"
	case "bool":
		return "bool"
	case "time":
		return "time.Time"
	case "uuid":
		return "uuid.UUID"
	default:
		return "string"
	}
}

// GetModuleName reads and extracts the module name from the go.mod file.
// Parses the go.mod file to determine the current project's module path,
// which is used for generating proper import statements in generated code.
func GetModuleName() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}
	return "", fmt.Errorf("module name not found in go.mod")
}

// GetCurrentTimestamp returns the current timestamp in a formatted string
func GetCurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(str string) string {
	parts := strings.Split(str, "_")
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

// ToPascalCase converts a string to PascalCase
func ToPascalCase(str string) string {
	parts := strings.Split(str, "_")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}