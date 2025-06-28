// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Handler templates
const HandlerTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package handler

import (
	"math"
	"net/http"

	"{{.ProjectName}}/dto/{{.PackageName}}"
	"{{.ProjectName}}/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// {{.ModelName}}Handler interface defines the contract for {{.ModelName}} HTTP handlers
type {{.ModelName}}Handler interface {
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type {{.VarName}}Handler struct {
	{{.VarName}}Service service.{{.ModelName}}Service
}

func New{{.ModelName}}Handler({{.VarName}}Service service.{{.ModelName}}Service) {{.ModelName}}Handler {
	return &{{.VarName}}Handler{
		{{.VarName}}Service: {{.VarName}}Service,
	}
}

// FindAll retrieves all {{.ModelName}}s with optional filtering and pagination
func (h *{{.VarName}}Handler) FindAll(ctx *fiber.Ctx) error {
	var filter {{.PackageName}}.Get{{.ModelName}}Dto
	
	// Parse query parameters
	_ = ctx.QueryParser(&filter)
	
	// Set default values
	filter.SetDefaults()
	
	// Get data from service
	{{.VarName}}s, total, err := h.{{.VarName}}Service.FindAll(ctx.Context(), &filter)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"message":   err.Error(),
		})
	}
	
	// Calculate pagination metadata
	lastPage := math.Ceil(float64(total) / float64(*filter.PageSize))
	
	return ctx.Status(http.StatusOK).JSON(map[string]any{
			"requestId": uuid.New(),
			"data":      {{.VarName}}s,
			"total":     total,
			"page":      filter.Page,
			"pageSize":  filter.PageSize,
			"lastPage":  lastPage,
		})
}

// FindById retrieves a single {{.ModelName}} by ID
func (h *{{.VarName}}Handler) FindById(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"message":   "Invalid id",
		})
	}
	
	{{.VarName}}, err := h.{{.VarName}}Service.FindById(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"message":   "{{.ModelName}} not found",
		})
	}
	
	return ctx.Status(http.StatusOK).JSON(map[string]any{
		"requestId": uuid.New(),
		"{{.VarName}}": {{.VarName}},
	})
}

// Create creates a new {{.ModelName}}
func (h *{{.VarName}}Handler) Create(ctx *fiber.Ctx) error {
	var request {{.PackageName}}.Create{{.ModelName}}Dto
	
	// Parse request body
	_ = ctx.BodyParser(&request)
	
	// Create via service
	{{.VarName}}, err := h.{{.VarName}}Service.Create(ctx.Context(), &request)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"{{.VarName}}": nil,
		})
	}
	
	return ctx.Status(http.StatusCreated).JSON(map[string]any{
		"requestId": uuid.New(),
		"{{.VarName}}": {{.VarName}},
	})
}

// Update updates an existing {{.ModelName}}
func (h *{{.VarName}}Handler) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"message":   "Invalid id",
		})
	}
	
	var request {{.PackageName}}.Update{{.ModelName}}Dto
	
	// Parse request body
	_ = ctx.BodyParser(&request)
	
	// Update via service
	err = h.{{.VarName}}Service.Update(ctx.Context(), id, &request)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"message":   "Something went wrong",
		})
	}
	
	return ctx.Status(http.StatusOK).JSON(map[string]any{
		"requestId": uuid.New(),
		"message":   "{{.ModelName}} updated successfully",
	})
}

// Delete deletes a {{.ModelName}} by ID
func (h *{{.VarName}}Handler) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"message":   "Invalid id",
		})
	}
	
	// Delete via service
	if err := h.{{.VarName}}Service.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"requestId": uuid.New(),
			"message":   "Something went wrong",
		})
	}
	
	return ctx.Status(http.StatusOK).JSON(map[string]any{
		"requestId": uuid.New(),
		"message":   "{{.ModelName}} deleted successfully",
	})
}
`