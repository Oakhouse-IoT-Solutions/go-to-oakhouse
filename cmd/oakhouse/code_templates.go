package main

import "strings"

// Model templates
const modelTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{.ModelName}} struct {
	ID        uuid.UUID      ` + "`gorm:\"type:uuid;primary_key;default:gen_random_uuid()\" json:\"id\"`" + `
{{range .Fields}}	{{.Name}} {{.Type}} ` + "`gorm:\"{{.GormTag}}\" json:\"{{.JsonTag}}\"`" + `
{{end}}	CreatedAt time.Time      ` + "`gorm:\"autoCreateTime\" json:\"created_at\"`" + `
	UpdatedAt time.Time      ` + "`gorm:\"autoUpdateTime\" json:\"updated_at\"`" + `
	DeletedAt gorm.DeletedAt ` + "`gorm:\"index\" json:\"deleted_at,omitempty\"`" + `
}

func ({{.ModelName}}) TableName() string {
	return "{{.TableName}}"
}
`

// Repository templates
const repositoryInterfaceTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package repository

import (
	"context"
	"{{.ProjectName}}/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{.ModelName}}Repository interface {
	FindAll(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]model.{{.ModelName}}, error)
	FindByID(ctx context.Context, id uuid.UUID, scopes ...func(*gorm.DB) *gorm.DB) (*model.{{.ModelName}}, error)
	Create(ctx context.Context, {{.VarName}} *model.{{.ModelName}}) error
	Update(ctx context.Context, {{.VarName}} *model.{{.ModelName}}) error
	Delete(ctx context.Context, id uuid.UUID) error
	Count(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error)
	FindWithPagination(ctx context.Context, offset, limit int, scopes ...func(*gorm.DB) *gorm.DB) ([]model.{{.ModelName}}, int64, error)
}
`

const repositoryImplTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package repository

import (
	"context"
	"{{.ProjectName}}/adapter"
	"{{.ProjectName}}/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{.VarName}}Repository struct {
	db *adapter.DatabaseAdapter
}

func New{{.ModelName}}Repository(db *adapter.DatabaseAdapter) {{.ModelName}}Repository {
	return &{{.VarName}}Repository{db: db}
}

func (r *{{.VarName}}Repository) FindAll(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]model.{{.ModelName}}, error) {
	var {{.VarName}}s []model.{{.ModelName}}
	query := r.db.DB.WithContext(ctx)
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.Find(&{{.VarName}}s).Error
	return {{.VarName}}s, err
}

func (r *{{.VarName}}Repository) FindByID(ctx context.Context, id uuid.UUID, scopes ...func(*gorm.DB) *gorm.DB) (*model.{{.ModelName}}, error) {
	var {{.VarName}} model.{{.ModelName}}
	query := r.db.DB.WithContext(ctx)
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.First(&{{.VarName}}, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &{{.VarName}}, nil
}

func (r *{{.VarName}}Repository) Create(ctx context.Context, {{.VarName}} *model.{{.ModelName}}) error {
	return r.db.DB.WithContext(ctx).Create({{.VarName}}).Error
}

func (r *{{.VarName}}Repository) Update(ctx context.Context, {{.VarName}} *model.{{.ModelName}}) error {
	return r.db.DB.WithContext(ctx).Save({{.VarName}}).Error
}

func (r *{{.VarName}}Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.DB.WithContext(ctx).Delete(&model.{{.ModelName}}{}, "id = ?", id).Error
}

func (r *{{.VarName}}Repository) Count(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	query := r.db.DB.WithContext(ctx).Model(&model.{{.ModelName}}{})
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.Count(&count).Error
	return count, err
}

func (r *{{.VarName}}Repository) FindWithPagination(ctx context.Context, offset, limit int, scopes ...func(*gorm.DB) *gorm.DB) ([]model.{{.ModelName}}, int64, error) {
	var {{.VarName}}s []model.{{.ModelName}}
	var total int64
	
	// Count total records
	countQuery := r.db.DB.WithContext(ctx).Model(&model.{{.ModelName}}{})
	for _, scope := range scopes {
		countQuery = scope(countQuery)
	}
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Get paginated records
	query := r.db.DB.WithContext(ctx).Offset(offset).Limit(limit)
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.Find(&{{.VarName}}s).Error
	return {{.VarName}}s, total, err
}
`

// Service templates
const serviceInterfaceTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package service

import (
	"context"
	"{{.ProjectName}}/dto/{{.PackageName}}"
	"{{.ProjectName}}/model"
	"github.com/google/uuid"
)

type {{.ModelName}}Service interface {
	FindAll(ctx context.Context, dto *{{.PackageName}}.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.{{.ModelName}}, error)
	Create(ctx context.Context, dto *{{.PackageName}}.Create{{.ModelName}}Dto) (*model.{{.ModelName}}, error)
	Update(ctx context.Context, id uuid.UUID, dto *{{.PackageName}}.Update{{.ModelName}}Dto) (*model.{{.ModelName}}, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindWithPagination(ctx context.Context, dto *{{.PackageName}}.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, int64, error)
}
`

const serviceImplTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package service

import (
	"context"
	"{{.ProjectName}}/dto/{{.PackageName}}"
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/repository"
	"{{.ProjectName}}/scope/{{.PackageName}}"
	"{{.ProjectName}}/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{.VarName}}Service struct {
	repo repository.{{.ModelName}}Repository
}

func New{{.ModelName}}Service(repo repository.{{.ModelName}}Repository) {{.ModelName}}Service {
	return &{{.VarName}}Service{repo: repo}
}

func (s *{{.VarName}}Service) FindAll(ctx context.Context, dto *{{.PackageName}}.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, error) {
	scopes := s.buildScopes(dto)
	return s.repo.FindAll(ctx, scopes...)
}

func (s *{{.VarName}}Service) FindByID(ctx context.Context, id uuid.UUID) (*model.{{.ModelName}}, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *{{.VarName}}Service) Create(ctx context.Context, dto *{{.PackageName}}.Create{{.ModelName}}Dto) (*model.{{.ModelName}}, error) {
	{{.VarName}} := &model.{{.ModelName}}{
{{range .Fields}}		{{.Name}}: dto.{{.Name}},
{{end}}	}
	
	if err := s.repo.Create(ctx, {{.VarName}}); err != nil {
		return nil, err
	}
	
	return {{.VarName}}, nil
}

func (s *{{.VarName}}Service) Update(ctx context.Context, id uuid.UUID, dto *{{.PackageName}}.Update{{.ModelName}}Dto) (*model.{{.ModelName}}, error) {
	{{.VarName}}, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
{{range .Fields}}	if dto.{{.Name}} != nil {
		{{$.VarName}}.{{.Name}} = *dto.{{.Name}}
	}
{{end}}	
	if err := s.repo.Update(ctx, {{.VarName}}); err != nil {
		return nil, err
	}
	
	return {{.VarName}}, nil
}

func (s *{{.VarName}}Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *{{.VarName}}Service) FindWithPagination(ctx context.Context, dto *{{.PackageName}}.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, int64, error) {
	scopes := s.buildScopes(dto)
	offset := (dto.Page - 1) * dto.PageSize
	return s.repo.FindWithPagination(ctx, offset, dto.PageSize, scopes...)
}

func (s *{{.VarName}}Service) buildScopes(dto *{{.PackageName}}.Get{{.ModelName}}Dto) []func(*gorm.DB) *gorm.DB {
	var scopes []func(*gorm.DB) *gorm.DB
	
	// Add your custom scopes here based on DTO fields
	// Example:
	// if dto.Status != "" {
	//     scopes = append(scopes, {{.PackageName}}.FilterByStatus(dto.Status))
	// }
	
	return scopes
}
`

// Handler templates
const handlerTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package handler

import (
	"github.com/gofiber/fiber/v2"
)

type {{.ModelName}}Handler struct {
}

func New{{.ModelName}}Handler() *{{.ModelName}}Handler {
	return &{{.ModelName}}Handler{}
}

func (h *{{.ModelName}}Handler) FindAll(c *fiber.Ctx) error {
	return c.SendString("All {{.ModelName}}s retrieved successfully")
}

func (h *{{.ModelName}}Handler) FindByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("{{.ModelName}} with ID " + idStr + " retrieved successfully")
}

func (h *{{.ModelName}}Handler) Create(c *fiber.Ctx) error {
	return c.SendString("{{.ModelName}} created successfully")
}

func (h *{{.ModelName}}Handler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("{{.ModelName}} with ID " + idStr + " updated successfully")
}

func (h *{{.ModelName}}Handler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("{{.ModelName}} with ID " + idStr + " deleted successfully")
}
`

// DTO templates
const getDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import "time"

type Get{{.ModelName}}Dto struct {
	Page     int ` + "`query:\"page\" validate:\"min=1\"`" + `
	PageSize int ` + "`query:\"page_size\" validate:\"min=1,max=100\"`" + `
	
	// Add your filter fields here
{{range .Fields}}	{{.Name}} {{.QueryType}} ` + "`query:\"{{.QueryTag}}\"`" + `
{{end}}	
	// Date range filters
	CreatedFrom *time.Time ` + "`query:\"created_from\"`" + `
	CreatedTo   *time.Time ` + "`query:\"created_to\"`" + `
}

func (dto *Get{{.ModelName}}Dto) SetDefaults() {
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
`

const createDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

type Create{{.ModelName}}Dto struct {
{{range .Fields}}	{{.Name}} {{.Type}} ` + "`json:\"{{.JsonTag}}\"`" + `
{{end}}
}
`

const updateDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

type Update{{.ModelName}}Dto struct {
{{range .Fields}}	{{.Name}} *{{.Type}} ` + "`json:\"{{.JsonTag}}\"`" + `
{{end}}
}
`

// Scope templates
const scopeTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import (
	"gorm.io/gorm"
)

// FilterBy{{.FieldName}} filters {{.ModelName}} by {{.FieldName}}
func FilterBy{{.FieldName}}({{.ParamName}} {{.ParamType}}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("{{.ColumnName}} = ?", {{.ParamName}})
	}
}

// FilterBy{{.FieldName}}In filters {{.ModelName}} by {{.FieldName}} in list
func FilterBy{{.FieldName}}In({{.ParamName}}s []{{.ParamType}}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len({{.ParamName}}s) == 0 {
			return db
		}
		return db.Where("{{.ColumnName}} IN ?", {{.ParamName}}s)
	}
}

// FilterBy{{.FieldName}}Like filters {{.ModelName}} by {{.FieldName}} with LIKE
func FilterBy{{.FieldName}}Like({{.ParamName}} string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if {{.ParamName}} == "" {
			return db
		}
		return db.Where("{{.ColumnName}} ILIKE ?", "%"+{{.ParamName}}+"%")
	}
}
`

// Middleware template
const middlewareTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package middleware

import (
	"{{.ProjectName}}/util"
	"github.com/gofiber/fiber/v2"
)

func {{.MiddlewareName}}() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Add your middleware logic here
		
		// Example: Check some condition
		// if someCondition {
		//     return util.SendError(c, fiber.StatusForbidden, "Access denied", nil)
		// }
		
		return c.Next()
	}
}
`

// Resource route template
const resourceRouteTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package route

import (
	"{{.ProjectName}}/handler"
	"github.com/gofiber/fiber/v2"
)

// Setup{{.Name}}Routes sets up routes for {{.Name}} resource
func Setup{{.Name}}Routes(api fiber.Router) {
	// Initialize handler
	{{.LowerName}}Handler := handler.New{{.Name}}Handler()

	// Setup routes
	{{.LowerName}}Group := api.Group("/{{.LowerName}}s")
	{{.LowerName}}Group.Get("/", {{.LowerName}}Handler.FindAll)
	{{.LowerName}}Group.Get("/:id", {{.LowerName}}Handler.FindByID)
	{{.LowerName}}Group.Post("/", {{.LowerName}}Handler.Create)
	{{.LowerName}}Group.Put("/:id", {{.LowerName}}Handler.Update)
	{{.LowerName}}Group.Delete("/:id", {{.LowerName}}Handler.Delete)
}
`

// Helper functions for template data
type TemplateData struct {
	ProjectName    string
	ModelName      string
	VarName        string
	PackageName    string
	TableName      string
	Fields         []FieldData
	MiddlewareName string
	FieldName      string
	ParamName      string
	ParamType      string
	ColumnName     string
}

type FieldData struct {
	Name        string
	Type        string
	QueryType   string
	GormTag     string
	JsonTag     string
	QueryTag    string
	ValidateTag string
}

// toSnakeCase converts CamelCase strings to snake_case format.
// Inserts underscores before uppercase letters and converts the entire string to lowercase,
// commonly used for database column names and file naming conventions.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

// toCamelCase converts strings to camelCase format by lowercasing the first character.
// Used for generating variable names and method parameters that follow Go naming conventions,
// ensuring consistent code style across generated files.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func toCamelCase(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToLower(str[:1]) + str[1:]
}

// toPlural converts singular nouns to their plural forms using basic English pluralization rules.
// Handles common cases like 'y' to 'ies', and adds 'es' or 's' as appropriate,
// used for generating table names and API endpoint paths.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func toPlural(str string) string {
	// Simple pluralization - you might want to use a proper library
	if strings.HasSuffix(str, "y") {
		return str[:len(str)-1] + "ies"
	}
	if strings.HasSuffix(str, "s") || strings.HasSuffix(str, "x") || strings.HasSuffix(str, "z") {
		return str + "es"
	}
	return str + "s"
}
