// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Scope templates
const ScopeTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import (
	"gorm.io/gorm"
)

// FilterBy{{.FieldName}} filters {{.ModelName}} by {{.FieldName}}
func FilterBy{{.FieldName}}({{.ParamName}} {{.ParamType}}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.ColumnName}} = ?", {{.ParamName}})
	}
}

// FilterBy{{.FieldName}}In filters {{.ModelName}} by {{.FieldName}} in list
func FilterBy{{.FieldName}}In({{.ParamName}}s []{{.ParamType}}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len({{.ParamName}}s) == 0 {
			return db
		}
		return db.Where("{{.ColumnName}} IN ?", {{.ParamName}}s)
	}
}

// FilterBy{{.FieldName}}Like filters {{.ModelName}} by {{.FieldName}} with LIKE
func FilterBy{{.FieldName}}Like({{.ParamName}} string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if {{.ParamName}} == "" {
			return db
		}
		return db.Where("{{.ColumnName}} ILIKE ?", "%"+{{.ParamName}}+"%")
	}
}
`