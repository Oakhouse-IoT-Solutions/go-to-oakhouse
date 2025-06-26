package repository

import (
	"context"
	"github.com/Oakhouse-Technology/go-to-oakhouse/adapter"
	"github.com/Oakhouse-Technology/go-to-oakhouse/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *adapter.DatabaseAdapter
}

func NewProductRepository(db *adapter.DatabaseAdapter) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]model.Product, error) {
	var products []model.Product
	query := r.db.DB.WithContext(ctx)
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(ctx context.Context, id uuid.UUID, scopes ...func(*gorm.DB) *gorm.DB) (*model.Product, error) {
	var product model.Product
	query := r.db.DB.WithContext(ctx)
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	return r.db.DB.WithContext(ctx).Create(product).Error
}

func (r *productRepository) Update(ctx context.Context, product *model.Product) error {
	return r.db.DB.WithContext(ctx).Save(product).Error
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.DB.WithContext(ctx).Delete(&model.Product{}, "id = ?", id).Error
}

func (r *productRepository) Count(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	query := r.db.DB.WithContext(ctx).Model(&model.Product{})
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.Count(&count).Error
	return count, err
}

func (r *productRepository) FindWithPagination(ctx context.Context, offset, limit int, scopes ...func(*gorm.DB) *gorm.DB) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64
	
	// Count total records
	countQuery := r.db.DB.WithContext(ctx).Model(&model.Product{})
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
	
	err := query.Find(&products).Error
	return products, total, err
}
