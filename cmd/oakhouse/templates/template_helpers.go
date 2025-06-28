// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

import "strings"

// Helper functions for template data
type TemplateData struct {
	ProjectName    string
	ModelName      string
	VarName        string
	PackageName    string
	TableName      string
	Fields         []FieldData
	MiddlewareName string
	FieldName      string
	ParamName      string
	ParamType      string
	ColumnName     string
}

type FieldData struct {
	Name        string
	Type        string
	QueryType   string
	GormTag     string
	JsonTag     string
	QueryTag    string
	ValidateTag string
}

// toSnakeCase converts CamelCase strings to snake_case format.
// Inserts underscores before uppercase letters and converts the entire string to lowercase,
// commonly used for database column names and file naming conventions.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
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

// toCamelCase converts strings to camelCase format by lowercasing the first character.
// Used for generating variable names and method parameters that follow Go naming conventions,
// ensuring consistent code style across generated files.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func ToCamelCase(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToLower(str[:1]) + str[1:]
}

// toPlural converts singular nouns to their plural forms using basic English pluralization rules.
// Handles common cases like 'y' to 'ies', and adds 'es' or 's' as appropriate,
// used for generating table names and API endpoint paths.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func ToPlural(str string) string {
	// Simple pluralization - you might want to use a proper library
	if strings.HasSuffix(str, "y") {
		return str[:len(str)-1] + "ies"
	}
	if strings.HasSuffix(str, "s") || strings.HasSuffix(str, "x") || strings.HasSuffix(str, "z") {
		return str + "es"
	}
	return str + "s"
}