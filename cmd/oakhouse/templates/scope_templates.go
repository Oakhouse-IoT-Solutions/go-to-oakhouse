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
		{{if eq .ParamType "string"}}return db.Where("{{.ColumnName}} ILIKE ?", "%"+{{.ParamName}}+"%"){{else}}return db.Where("{{.ColumnName}} = ?", {{.ParamName}}){{end}}
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

// DateRangeFilterTemplate for date range filtering
const DateRangeFilterTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import (
	"time"
	"gorm.io/gorm"
)

// FilterByDateRange filters {{.ModelName}} by {{.FieldName}} date range
func FilterByDateRange(startDate, endDate time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !startDate.IsZero() && !endDate.IsZero() {
			return db.Where("{{.ColumnName}} BETWEEN ? AND ?", startDate, endDate)
		} else if !startDate.IsZero() {
			return db.Where("{{.ColumnName}} >= ?", startDate)
		} else if !endDate.IsZero() {
			return db.Where("{{.ColumnName}} <= ?", endDate)
		}
		return db
	}
}
`

// AdvancedDateRangeFilterTemplate for more complex date range filtering
const AdvancedDateRangeFilterTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import (
	"time"
	"gorm.io/gorm"
)

// DateRangeFilter applies date range filtering to {{.ModelName}}
type DateRangeFilter struct {
	StartDate *time.Time ` + "`json:\"start_date\" query:\"start_date\"`" + `
	EndDate   *time.Time ` + "`json:\"end_date\" query:\"end_date\"`" + `
}

// Apply applies the date range filter
func (f *DateRangeFilter) Apply(db *gorm.DB, column string) *gorm.DB {
	if f.StartDate != nil {
		db = db.Where(column+" >= ?", *f.StartDate)
	}
	if f.EndDate != nil {
		db = db.Where(column+" <= ?", *f.EndDate)
	}
	return db
}
`

// PaginationScopeTemplate for pagination
const PaginationScopeTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import (
	"gorm.io/gorm"
)

// PaginationScope applies pagination to {{.ModelName}} queries
func PaginationScope(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// CountScope returns total count for {{.ModelName}}
func CountScope(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Count(&count).Error
	return count, err
}
`