// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package ticket2

import (
	"time"
	"gorm.io/gorm"
)

// FilterByDateRange filters Ticket2 by created_at date range
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

// FilterByTitle filters Ticket2 by title
func FilterByTitle(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title ILIKE ?", "%"+title+"%")
	}
}

// FilterByDescription filters Ticket2 by description
func FilterByDescription(description string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("description ILIKE ?", "%"+description+"%")
	}
}
