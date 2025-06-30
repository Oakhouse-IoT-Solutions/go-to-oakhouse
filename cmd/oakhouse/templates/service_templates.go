// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Service templates
const ServiceInterfaceTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package service

import (
	"context"
	"{{.ProjectName}}/dto/{{.PackageName}}"
	"{{.ProjectName}}/model"
	"github.com/google/uuid"
)

// {{.ModelName}}Service defines the interface for {{.PackageName}}-related operations
type {{.ModelName}}Service interface {
	// FindAll retrieves all {{.PackageName}}s with optional filtering
	FindAll(ctx context.Context, dto *{{.PackageName}}.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, int64, error)
	
	// FindById retrieves a {{.PackageName}} by its ID
	FindById(ctx context.Context, id uuid.UUID) (*model.{{.ModelName}}, error)
	
	// Create creates a new {{.PackageName}}
	Create(ctx context.Context, dto *{{.PackageName}}.Create{{.ModelName}}Dto) (*model.{{.ModelName}}, error)
	
	// Update updates an existing {{.PackageName}}
	Update(ctx context.Context, id uuid.UUID, dto *{{.PackageName}}.Update{{.ModelName}}Dto) error
	
	// Delete removes a {{.PackageName}} by its ID
	Delete(ctx context.Context, id uuid.UUID) error
}
`

const ServiceTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	dto "{{.ProjectName}}/dto/{{.PackageName}}"
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/repository"
	tscope "{{.ProjectName}}/scope/{{.PackageName}}"
)

type {{.VarName}}Service struct {
	repo repository.{{.ModelName}}Repository
}

func New{{.ModelName}}Service(repo repository.{{.ModelName}}Repository) {{.ModelName}}Service {
	return &{{.VarName}}Service{repo: repo}
}

func (s *{{.VarName}}Service) FindAll(ctx context.Context, getDto *dto.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, int64, error) {
	scopes := s.buildScopes(getDto)
	offset := (*getDto.Page - 1) * *getDto.PageSize
	return s.repo.FindWithPagination(ctx, offset, *getDto.PageSize, scopes...)
}

func (s *{{.VarName}}Service) FindById(ctx context.Context, id uuid.UUID) (*model.{{.ModelName}}, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *{{.VarName}}Service) Create(ctx context.Context, createDto *dto.Create{{.ModelName}}Dto) (*model.{{.ModelName}}, error) {
	new{{.ModelName}} := &model.{{.ModelName}}{
{{range .Fields}}		{{.Name}}: createDto.{{.Name}},
{{end}}	}
	
	if err := s.repo.Create(ctx, new{{.ModelName}}); err != nil {
		return nil, err
	}
	
	return new{{.ModelName}}, nil
}

func (s *{{.VarName}}Service) Update(ctx context.Context, id uuid.UUID, updateDto *dto.Update{{.ModelName}}Dto) error {
	existing{{.ModelName}}, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	
{{range .Fields}}	if updateDto.{{.Name}} != nil {
		existing{{$.ModelName}}.{{.Name}} = *updateDto.{{.Name}}
	}
{{end}}	
	return s.repo.Update(ctx, existing{{.ModelName}})
}

func (s *{{.VarName}}Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *{{.VarName}}Service) buildScopes(getDto *dto.Get{{.ModelName}}Dto) []func(*gorm.DB) *gorm.DB {
	var scopes []func(*gorm.DB) *gorm.DB
	
	// Add field-specific filters
{{range .Fields}}{{if eq .Type "string"}}	if getDto.{{.Name}} != nil && *getDto.{{.Name}} != "" {
		scopes = append(scopes, tscope.FilterBy{{.Name}}(*getDto.{{.Name}}))
	}
{{else}}	if getDto.{{.Name}} != nil {
		scopes = append(scopes, tscope.FilterBy{{.Name}}(*getDto.{{.Name}}))
	}
{{end}}{{end}}	
	// Add date range filtering if available
	if getDto.StartDate != nil || getDto.EndDate != nil {
		var startTime, endTime time.Time
		if getDto.StartDate != nil {
			startTime = *getDto.StartDate
		}
		if getDto.EndDate != nil {
			endTime = *getDto.EndDate
		}
		scopes = append(scopes, tscope.FilterByDateRange(startTime, endTime))
	}
	
	return scopes
}
`

// Simple service implementation template for rapid prototyping without database dependencies
const SimpleServiceTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package service

// {{.ModelName}}Service implements basic operations without database dependencies
type {{.ModelName}}Service struct{}

func New{{.ModelName}}Service() {{.ModelName}}Service {
	return {{.ModelName}}Service{}
}

func (s {{.ModelName}}Service) Index() string {
	return "{{.ModelName}}s working"
}
`
