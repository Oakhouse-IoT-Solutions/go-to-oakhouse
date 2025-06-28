// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package tickeeeet

import (
	"time"
	"gorm.io/gorm"
)

// FilterByTitle filters Tickeeeet by title using ILIKE for case-insensitive search
func FilterByTitle(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title ILIKE ?", "%"+title+"%")
	}
}

// FilterByDescription filters Tickeeeet by description using ILIKE for case-insensitive search
func FilterByDescription(description string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("description ILIKE ?", "%"+description+"%")
	}
}

// FilterByDateRange filters Tickeeeet by created_at date range
func FilterByDateRange(startDate, endDate time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !startDate.IsZero() && !endDate.IsZero() {
			return db.Where("created_at BETWEEN ? AND ?", startDate, endDate)
		} else if !startDate.IsZero() {
			return db.Where("created_at >= ?", startDate)
		} else if !endDate.IsZero() {
			return db.Where("created_at <= ?", endDate)
		}
		return db
	}
}
