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

type {{.ModelName}}Service interface {
	FindAll(ctx context.Context, dto *{{.PackageName}}.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.{{.ModelName}}, error)
	Create(ctx context.Context, dto *{{.PackageName}}.Create{{.ModelName}}Dto) (*model.{{.ModelName}}, error)
	Update(ctx context.Context, id uuid.UUID, dto *{{.PackageName}}.Update{{.ModelName}}Dto) (*model.{{.ModelName}}, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindWithPagination(ctx context.Context, dto *{{.PackageName}}.Get{{.ModelName}}Dto) ([]model.{{.ModelName}}, int64, error)
}
`

const ServiceImplTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
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
	
	// Add pagination scope
	scopes = append(scopes, {{.PackageName}}.PaginationScope(dto.Page, dto.PageSize))
	
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
	
	// Add date range filter
	scopes = append(scopes, {{.PackageName}}.DateRangeFilter(dto))
	
	// Add your custom scopes here based on DTO fields
	// Example:
	// if dto.Status != \"\" {
	//     scopes = append(scopes, {{.PackageName}}.FilterByStatus(dto.Status))
	// }
	
	return scopes
}
`