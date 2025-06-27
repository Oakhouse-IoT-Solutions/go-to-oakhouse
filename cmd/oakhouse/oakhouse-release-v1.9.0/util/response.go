package util

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}

func PaginatedSuccessResponse(message string, data interface{}, pagination Pagination) PaginatedResponse {
	return PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}

func SendSuccess(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(SuccessResponse(message, data))
}

func SendError(c *fiber.Ctx, status int, message string, err interface{}) error {
	return c.Status(status).JSON(ErrorResponse(message, err))
}

func SendPaginatedSuccess(c *fiber.Ctx, message string, data interface{}, pagination Pagination) error {
	return c.JSON(PaginatedSuccessResponse(message, data, pagination))
}
