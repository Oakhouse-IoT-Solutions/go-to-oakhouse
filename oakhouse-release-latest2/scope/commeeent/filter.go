// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commeeent

import (
	"time"
	"gorm.io/gorm"
)

// FilterByDateRange filters Commeeent by DateRange date range
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

// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡

// FilterByTitle filters Commeeent by Title
func FilterByTitle(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title ILIKE ?", "%"+title+"%")
	}
}

// FilterByTitleIn filters Commeeent by Title in list
func FilterByTitleIn(titles []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(titles) == 0 {
			return db
		}
		return db.Where("title IN ?", titles)
	}
}

// FilterByTitleLike filters Commeeent by Title with LIKE
func FilterByTitleLike(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if title == "" {
			return db
		}
		return db.Where("title ILIKE ?", "%"+title+"%")
	}
}

// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡

// FilterByDescription filters Commeeent by Description
func FilterByDescription(description string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("description ILIKE ?", "%"+description+"%")
	}
}

// FilterByDescriptionIn filters Commeeent by Description in list
func FilterByDescriptionIn(descriptions []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(descriptions) == 0 {
			return db
		}
		return db.Where("description IN ?", descriptions)
	}
}

// FilterByDescriptionLike filters Commeeent by Description with LIKE
func FilterByDescriptionLike(description string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if description == "" {
			return db
		}
		return db.Where("description ILIKE ?", "%"+description+"%")
	}
}
