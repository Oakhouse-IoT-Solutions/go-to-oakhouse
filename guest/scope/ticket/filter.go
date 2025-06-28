// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package ticket

import (
	"gorm.io/gorm"
)

// FilterByfilter filters Ticket by filter
func FilterByfilter(<no value> <no value>) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("<no value> = ?", <no value>)
	}
}

// FilterByfilterIn filters Ticket by filter in list
func FilterByfilterIn(<no value>s []<no value>) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(<no value>s) == 0 {
			return db
		}
		return db.Where("<no value> IN ?", <no value>s)
	}
}

// FilterByfilterLike filters Ticket by filter with LIKE
func FilterByfilterLike(<no value> string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if <no value> == "" {
			return db
		}
		return db.Where("<no value> ILIKE ?", "%"+<no value>+"%")
	}
}
