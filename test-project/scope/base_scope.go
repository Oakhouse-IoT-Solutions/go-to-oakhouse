// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package scope

import (
	"time"

	"gorm.io/gorm"
)

// DateRangeScope applies date range filtering with pointer types for consistency
func DateRangeScope(startDate, endDate *time.Time, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startDate != nil {
			db = db.Where(column+" >= ?", *startDate)
		}
		if endDate != nil {
			db = db.Where(column+" <= ?", *endDate)
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
