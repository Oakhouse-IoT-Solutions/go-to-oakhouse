// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package user

import (
	"time"
	"gorm.io/gorm"
)

// FilterByDateRange filters User by DateRange date range
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

// FilterByName filters User by Name
func FilterByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name ILIKE ?", "%"+name+"%")
	}
}

// FilterByNameIn filters User by Name in list
func FilterByNameIn(names []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) == 0 {
			return db
		}
		return db.Where("name IN ?", names)
	}
}

// FilterByNameLike filters User by Name with LIKE
func FilterByNameLike(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name == "" {
			return db
		}
		return db.Where("name ILIKE ?", "%"+name+"%")
	}
}

// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡

// FilterByEmail filters User by Email
func FilterByEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("email ILIKE ?", "%"+email+"%")
	}
}

// FilterByEmailIn filters User by Email in list
func FilterByEmailIn(emails []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(emails) == 0 {
			return db
		}
		return db.Where("email IN ?", emails)
	}
}

// FilterByEmailLike filters User by Email with LIKE
func FilterByEmailLike(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if email == "" {
			return db
		}
		return db.Where("email ILIKE ?", "%"+email+"%")
	}
}

// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡

// FilterByAge filters User by Age
func FilterByAge(age int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("age = ?", age)
	}
}

// FilterByAgeIn filters User by Age in list
func FilterByAgeIn(ages []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ages) == 0 {
			return db
		}
		return db.Where("age IN ?", ages)
	}
}

// FilterByAgeLike filters User by Age with LIKE
func FilterByAgeLike(age string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if age == "" {
			return db
		}
		return db.Where("age ILIKE ?", "%"+age+"%")
	}
}
