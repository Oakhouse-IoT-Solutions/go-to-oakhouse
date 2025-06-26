package service

import (
	"context"
	"github.com/Oakhouse-Technology/go-to-oakhouse/dto/product"
	"github.com/Oakhouse-Technology/go-to-oakhouse/model"
	"github.com/Oakhouse-Technology/go-to-oakhouse/repository"
	"github.com/Oakhouse-Technology/go-to-oakhouse/scope/product"
	"github.com/Oakhouse-Technology/go-to-oakhouse/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) FindAll(ctx context.Context, dto *product.GetProductDto) ([]model.Product, error) {
	scopes := s.buildScopes(dto)
	return s.repo.FindAll(ctx, scopes...)
}

func (s *productService) FindByID(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *productService) Create(ctx context.Context, dto *product.CreateProductDto) (*model.Product, error) {
	product := &model.Product{
	}
	
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}
	
	return product, nil
}

func (s *productService) Update(ctx context.Context, id uuid.UUID, dto *product.UpdateProductDto) (*model.Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	
	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}
	
	return product, nil
}

func (s *productService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *productService) FindWithPagination(ctx context.Context, dto *product.GetProductDto) ([]model.Product, int64, error) {
	scopes := s.buildScopes(dto)
	offset := (dto.Page - 1) * dto.PageSize
	return s.repo.FindWithPagination(ctx, offset, dto.PageSize, scopes...)
}

func (s *productService) buildScopes(dto *product.GetProductDto) []func(*gorm.DB) *gorm.DB {
	var scopes []func(*gorm.DB) *gorm.DB
	
	// Add your custom scopes here based on DTO fields
	// Example:
	// if dto.Status != "" {
	//     scopes = append(scopes, product.FilterByStatus(dto.Status))
	// }
	
	return scopes
}
