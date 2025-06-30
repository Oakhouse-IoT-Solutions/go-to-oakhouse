// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import "math"

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
	HasNext     bool  `json:"has_next"`
	HasPrev     bool  `json:"has_prev"`
}

// CalculatePagination calculates pagination metadata
func CalculatePagination(page, pageSize int, totalItems int64) PaginationMeta {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	return PaginationMeta{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
	}
}

// GetOffset calculates the database offset for pagination
func GetOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}
