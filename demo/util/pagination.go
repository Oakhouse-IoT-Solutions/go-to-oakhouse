// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PaginationParams struct {
	Page     int
	PageSize int
	Offset   int
}

func GetPaginationParams(c *fiber.Ctx) PaginationParams {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
	}
}

func CalculatePagination(page, pageSize int, total int64) Pagination {
	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))

	return Pagination{
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: totalPage,
	}
}
