package repository

import (
	"context"
	"github.com/Oakhouse-Technology/go-to-oakhouse/adapter"
	"github.com/Oakhouse-Technology/go-to-oakhouse/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *adapter.DatabaseAdapter
}

func NewUserRepository(db *adapter.DatabaseAdapter) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]model.User, error) {
	var users []model.User
	query := r.db.DB.WithContext(ctx)
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID, scopes ...func(*gorm.DB) *gorm.DB) (*model.User, error) {
	var user model.User
	query := r.db.DB.WithContext(ctx)
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.DB.WithContext(ctx).Create(user).Error
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.DB.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.DB.WithContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}

func (r *userRepository) Count(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	query := r.db.DB.WithContext(ctx).Model(&model.User{})
	
	for _, scope := range scopes {
		query = scope(query)
	}
	
	err := query.Count(&count).Error
	return count, err
}

func (r *userRepository) FindWithPagination(ctx context.Context, offset, limit int, scopes ...func(*gorm.DB) *gorm.DB) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	
	// Count total records
	countQuery := r.db.DB.WithContext(ctx).Model(&model.User{})
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
	
	err := query.Find(&users).Error
	return users, total, err
}
