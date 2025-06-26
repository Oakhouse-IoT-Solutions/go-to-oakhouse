package product

import "time"

type GetProductDto struct {
	Page     int `query:"page" validate:"min=1"`
	PageSize int `query:"page_size" validate:"min=1,max=100"`
	
	// Add your filter fields here
	
	// Date range filters
	CreatedFrom *time.Time `query:"created_from"`
	CreatedTo   *time.Time `query:"created_to"`
}

func (dto *GetProductDto) SetDefaults() {
	if dto.Page <= 0 {
		dto.Page = 1
	}
	if dto.PageSize <= 0 {
		dto.PageSize = 10
	}
	if dto.PageSize > 100 {
		dto.PageSize = 100
	}
}
