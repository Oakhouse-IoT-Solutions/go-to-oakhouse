// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Repository templates
const RepositoryInterfaceTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
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

const RepositoryImplTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
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
	
	err := query.First(&{{.VarName}}, \"id = ?\", id).Error
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
	return r.db.DB.WithContext(ctx).Delete(&model.{{.ModelName}}{}, \"id = ?\", id).Error
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