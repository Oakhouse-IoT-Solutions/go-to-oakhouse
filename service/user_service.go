package service

import (
	"context"
	"github.com/Oakhouse-Technology/go-to-oakhouse/dto/user"
	"github.com/Oakhouse-Technology/go-to-oakhouse/model"
	"github.com/Oakhouse-Technology/go-to-oakhouse/repository"
	"github.com/Oakhouse-Technology/go-to-oakhouse/scope/user"
	"github.com/Oakhouse-Technology/go-to-oakhouse/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindAll(ctx context.Context, dto *user.GetUserDto) ([]model.User, error) {
	scopes := s.buildScopes(dto)
	return s.repo.FindAll(ctx, scopes...)
}

func (s *userService) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) Create(ctx context.Context, dto *user.CreateUserDto) (*model.User, error) {
	user := &model.User{
	}
	
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	
	return user, nil
}

func (s *userService) Update(ctx context.Context, id uuid.UUID, dto *user.UpdateUserDto) (*model.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	
	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	
	return user, nil
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) FindWithPagination(ctx context.Context, dto *user.GetUserDto) ([]model.User, int64, error) {
	scopes := s.buildScopes(dto)
	offset := (dto.Page - 1) * dto.PageSize
	return s.repo.FindWithPagination(ctx, offset, dto.PageSize, scopes...)
}

func (s *userService) buildScopes(dto *user.GetUserDto) []func(*gorm.DB) *gorm.DB {
	var scopes []func(*gorm.DB) *gorm.DB
	
	// Add your custom scopes here based on DTO fields
	// Example:
	// if dto.Status != "" {
	//     scopes = append(scopes, user.FilterByStatus(dto.Status))
	// }
	
	return scopes
}
