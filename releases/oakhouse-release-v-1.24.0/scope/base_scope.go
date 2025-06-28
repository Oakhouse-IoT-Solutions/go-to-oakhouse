// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package scope

import "gorm.io/gorm"

// Omit excludes specified fields from SELECT
func Omit(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(fields...)
	}
}

// Select includes only specified fields in SELECT
func Select(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields...)
	}
}

// Unscoped includes soft deleted records
func Unscoped() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}
}

// OrderBy adds ORDER BY clause
func OrderBy(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// Limit adds LIMIT clause
func Limit(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

// Offset adds OFFSET clause
func Offset(offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}
