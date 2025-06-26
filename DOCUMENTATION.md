# Go To Oakhouse Framework Documentation

## Table of Contents

1. [Installation](#installation)
2. [Quick Start](#quick-start)
3. [Project Structure](#project-structure)
4. [CLI Commands](#cli-commands)
5. [Configuration](#configuration)
6. [Models and Entities](#models-and-entities)
7. [Repositories](#repositories)
8. [Services](#services)
9. [Handlers](#handlers)
10. [DTOs (Data Transfer Objects)](#dtos-data-transfer-objects)
11. [Scopes](#scopes)
12. [Middleware](#middleware)
13. [Database Operations](#database-operations)
14. [Authentication](#authentication)
15. [Testing](#testing)
16. [Deployment](#deployment)
17. [Best Practices](#best-practices)
18. [Examples](#examples)

## Installation

### Prerequisites

- Go 1.21 or higher
- PostgreSQL (recommended) or MySQL
- Redis (optional, for caching)

### Install CLI Tool

```bash
go install github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse@latest
```

### Verify Installation

```bash
oakhouse --version
```

## Quick Start

### Create a New Project

```bash
# Create a new project
oakhouse new my-blog-api
cd my-blog-api

# Install dependencies
go mod tidy

# Set up environment
cp .env.example .env
# Edit .env with your database credentials

# Run migrations
oakhouse migrate up

# Start development server
oakhouse serve
```

### Generate Your First Resource

```bash
# Generate a complete blog post resource
oakhouse generate resource Post title:string content:text author_id:uuid published:bool

# This creates:
# - entity/post.go (model)
# - repository/post_repository.go (data access)
# - service/post_service.go (business logic)
# - handler/post_handler.go (HTTP handlers)
# - dto/post/ (DTOs for API)
# - scope/post/ (query scopes)
```

## Project Structure

```
my-blog-api/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── config.go              # Configuration management
├── entity/
│   ├── user.go                # User model
│   └── post.go                # Post model
├── repository/
│   ├── user_repository.go     # User data access
│   └── post_repository.go     # Post data access
├── service/
│   ├── user_service.go        # User business logic
│   └── post_service.go        # Post business logic
├── handler/
│   ├── user_handler.go        # User HTTP handlers
│   └── post_handler.go        # Post HTTP handlers
├── dto/
│   ├── user/
│   │   ├── create_user_dto.go
│   │   ├── update_user_dto.go
│   │   └── get_user_dto.go
│   └── post/
│       ├── create_post_dto.go
│       ├── update_post_dto.go
│       └── get_post_dto.go
├── scope/
│   ├── user/
│   │   └── filter_by_role.go
│   └── post/
│       ├── filter_by_author.go
│       └── filter_by_status.go
├── middleware/
│   ├── auth.go                # Authentication middleware
│   └── rate_limit.go          # Rate limiting
├── adapter/
│   ├── database.go            # Database adapter
│   └── postgres/
│       └── connection.go      # PostgreSQL connection
├── route/
│   └── v1.go                  # API routes
├── util/
│   ├── response.go            # Response utilities
│   └── pagination.go          # Pagination utilities
├── migrations/
│   └── 001_create_users.sql   # Database migrations
├── .env.example               # Environment template
├── .env                       # Environment variables
├── docker-compose.yml         # Docker services
├── Dockerfile                 # Application container
├── Makefile                   # Build commands
└── go.mod                     # Go dependencies
```

## CLI Commands

### Project Management

```bash
# Create new project
oakhouse new <project-name>

# Build application
oakhouse build

# Start development server with hot reload
oakhouse serve

# Start production server
oakhouse serve --env=production
```

### Code Generation

```bash
# Generate complete resource (model, repository, service, handler, DTOs)
oakhouse generate resource User name:string email:string age:int

# Generate individual components
oakhouse generate model Product
oakhouse generate repository ProductRepository
oakhouse generate service ProductService
oakhouse generate handler ProductHandler
oakhouse generate dto product CreateProductDto
oakhouse generate scope product FilterByCategory
oakhouse generate middleware RoleCheck
```

### Database Operations

```bash
# Run all pending migrations
oakhouse migrate up

# Rollback last migration
oakhouse migrate down

# Create new migration
oakhouse migrate create add_posts_table

# Check migration status
oakhouse migrate status
```

## Configuration

### Environment Variables

Create a `.env` file in your project root:

```env
# Application
APP_NAME=my-blog-api
APP_PORT=8080
APP_ENV=development
APP_DEBUG=true

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=blog_db
DB_USER=postgres
DB_PASSWORD=password
DB_SSL_MODE=disable
DB_TIMEZONE=UTC

# JWT
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRES_IN=24h

# Redis
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# CORS
CORS_ALLOWED_ORIGINS=*
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=*
```

### Configuration Structure

```go
// config/config.go
package config

type Config struct {
    AppName  string `env:"APP_NAME" default:"go-to-oakhouse"`
    AppPort  string `env:"APP_PORT" default:"8080"`
    AppEnv   string `env:"APP_ENV" default:"development"`
    AppDebug bool   `env:"APP_DEBUG" default:"true"`
    
    DBHost     string `env:"DB_HOST" default:"localhost"`
    DBPort     string `env:"DB_PORT" default:"5432"`
    DBName     string `env:"DB_NAME" default:"oakhouse_db"`
    DBUser     string `env:"DB_USER" default:"postgres"`
    DBPassword string `env:"DB_PASSWORD" default:"password"`
}

func Load() *Config {
    // Implementation loads from environment
}
```

## Models and Entities

### Basic Model

```go
// entity/user.go
package entity

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Name      string         `gorm:"not null" json:"name" validate:"required,min=2,max=100"`
    Email     string         `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
    Password  string         `gorm:"not null" json:"-" validate:"required,min=8"`
    Role      string         `gorm:"default:user" json:"role"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    
    // Relationships
    Posts []Post `gorm:"foreignKey:AuthorID" json:"posts,omitempty"`
}

func (User) TableName() string {
    return "users"
}
```

### Model with Relationships

```go
// entity/post.go
package entity

type Post struct {
    ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Title       string         `gorm:"not null" json:"title" validate:"required,min=5,max=200"`
    Content     string         `gorm:"type:text" json:"content" validate:"required,min=10"`
    AuthorID    uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`
    Published   bool           `gorm:"default:false" json:"published"`
    PublishedAt *time.Time     `json:"published_at,omitempty"`
    CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    
    // Relationships
    Author User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
}

func (Post) TableName() string {
    return "posts"
}
```

## Repositories

### Repository Interface

```go
// repository/user_repository.go
package repository

import (
    "context"
    "my-blog-api/entity"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type UserRepository interface {
    FindAll(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]entity.User, error)
    FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
    Create(ctx context.Context, user *entity.User) error
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id uuid.UUID) error
    Count(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error)
    FindWithPagination(ctx context.Context, offset, limit int, scopes ...func(*gorm.DB) *gorm.DB) ([]entity.User, int64, error)
}
```

### Repository Implementation

```go
type userRepository struct {
    db *adapter.DatabaseAdapter
}

func NewUserRepository(db *adapter.DatabaseAdapter) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) FindAll(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]entity.User, error) {
    var users []entity.User
    query := r.db.DB.WithContext(ctx)
    
    for _, scope := range scopes {
        query = scope(query)
    }
    
    err := query.Find(&users).Error
    return users, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
    var user entity.User
    err := r.db.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

## Services

### Service Interface

```go
// service/user_service.go
package service

import (
    "context"
    "my-blog-api/dto/user"
    "my-blog-api/entity"
    "github.com/google/uuid"
)

type UserService interface {
    GetUsers(ctx context.Context, dto *user.GetUserDto) ([]entity.User, int64, error)
    GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
    CreateUser(ctx context.Context, dto *user.CreateUserDto) (*entity.User, error)
    UpdateUser(ctx context.Context, id uuid.UUID, dto *user.UpdateUserDto) (*entity.User, error)
    DeleteUser(ctx context.Context, id uuid.UUID) error
    AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error)
}
```

### Service Implementation

```go
type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, dto *user.CreateUserDto) (*entity.User, error) {
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    
    user := &entity.User{
        Name:     dto.Name,
        Email:    dto.Email,
        Password: string(hashedPassword),
        Role:     "user",
        IsActive: true,
    }
    
    if err := s.repo.Create(ctx, user); err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *userService) AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error) {
    user, err := s.repo.FindByEmail(ctx, email)
    if err != nil {
        return nil, err
    }
    
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid credentials")
    }
    
    return user, nil
}
```

## Handlers

### Handler Implementation

```go
// handler/user_handler.go
package handler

import (
    "my-blog-api/dto/user"
    "my-blog-api/service"
    "my-blog-api/util"
    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type UserHandler struct {
    service   service.UserService
    validator *validator.Validate
}

func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{
        service:   service,
        validator: validator.New(),
    }
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
    var dto user.GetUserDto
    if err := c.QueryParser(&dto); err != nil {
        return util.SendError(c, fiber.StatusBadRequest, "Invalid query parameters", err.Error())
    }
    
    dto.SetDefaults()
    
    if err := h.validator.Struct(&dto); err != nil {
        return util.SendError(c, fiber.StatusBadRequest, "Validation failed", err.Error())
    }
    
    users, total, err := h.service.GetUsers(c.Context(), &dto)
    if err != nil {
        return util.SendError(c, fiber.StatusInternalServerError, "Failed to fetch users", err.Error())
    }
    
    pagination := util.CalculatePagination(dto.Page, dto.PageSize, total)
    return util.SendPaginatedSuccess(c, "Users retrieved successfully", users, pagination)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var dto user.CreateUserDto
    if err := c.BodyParser(&dto); err != nil {
        return util.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
    }
    
    if err := h.validator.Struct(&dto); err != nil {
        return util.SendError(c, fiber.StatusBadRequest, "Validation failed", err.Error())
    }
    
    user, err := h.service.CreateUser(c.Context(), &dto)
    if err != nil {
        return util.SendError(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
    }
    
    return util.SendSuccess(c, "User created successfully", user)
}
```

## DTOs (Data Transfer Objects)

### Input DTOs

```go
// dto/user/create_user_dto.go
package user

type CreateUserDto struct {
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    Role     string `json:"role" validate:"omitempty,oneof=admin user moderator"`
}

// dto/user/update_user_dto.go
type UpdateUserDto struct {
    Name     *string `json:"name" validate:"omitempty,min=2,max=100"`
    Email    *string `json:"email" validate:"omitempty,email"`
    Role     *string `json:"role" validate:"omitempty,oneof=admin user moderator"`
    IsActive *bool   `json:"is_active"`
}
```

### Query DTOs

```go
// dto/user/get_user_dto.go
package user

import "time"

type GetUserDto struct {
    Page     int `query:"page" validate:"min=1"`
    PageSize int `query:"page_size" validate:"min=1,max=100"`
    
    // Filters
    Name     string `query:"name"`
    Email    string `query:"email"`
    Role     string `query:"role"`
    IsActive *bool  `query:"is_active"`
    
    // Date filters
    CreatedFrom *time.Time `query:"created_from"`
    CreatedTo   *time.Time `query:"created_to"`
    
    // Sorting
    SortBy    string `query:"sort_by" validate:"omitempty,oneof=name email created_at updated_at"`
    SortOrder string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

func (dto *GetUserDto) SetDefaults() {
    if dto.Page <= 0 {
        dto.Page = 1
    }
    if dto.PageSize <= 0 {
        dto.PageSize = 10
    }
    if dto.PageSize > 100 {
        dto.PageSize = 100
    }
    if dto.SortBy == "" {
        dto.SortBy = "created_at"
    }
    if dto.SortOrder == "" {
        dto.SortOrder = "desc"
    }
}
```

## Scopes

Scopes are reusable query functions that can be applied to GORM queries:

```go
// scope/user/filter_by_role.go
package user

import "gorm.io/gorm"

func FilterByRole(role string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if role == "" {
            return db
        }
        return db.Where("role = ?", role)
    }
}

func FilterByActive(isActive bool) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("is_active = ?", isActive)
    }
}

func FilterByNameLike(name string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if name == "" {
            return db
        }
        return db.Where("name ILIKE ?", "%"+name+"%")
    }
}

func FilterByDateRange(from, to *time.Time) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if from != nil {
            db = db.Where("created_at >= ?", from)
        }
        if to != nil {
            db = db.Where("created_at <= ?", to)
        }
        return db
    }
}
```

### Using Scopes in Services

```go
func (s *userService) GetUsers(ctx context.Context, dto *user.GetUserDto) ([]entity.User, int64, error) {
    var scopes []func(*gorm.DB) *gorm.DB
    
    // Apply filters based on DTO
    if dto.Role != "" {
        scopes = append(scopes, userScope.FilterByRole(dto.Role))
    }
    
    if dto.IsActive != nil {
        scopes = append(scopes, userScope.FilterByActive(*dto.IsActive))
    }
    
    if dto.Name != "" {
        scopes = append(scopes, userScope.FilterByNameLike(dto.Name))
    }
    
    if dto.CreatedFrom != nil || dto.CreatedTo != nil {
        scopes = append(scopes, userScope.FilterByDateRange(dto.CreatedFrom, dto.CreatedTo))
    }
    
    // Add sorting
    orderBy := fmt.Sprintf("%s %s", dto.SortBy, dto.SortOrder)
    scopes = append(scopes, scope.OrderBy(orderBy))
    
    offset := (dto.Page - 1) * dto.PageSize
    return s.repo.FindWithPagination(ctx, offset, dto.PageSize, scopes...)
}
```

## Middleware

### Authentication Middleware

```go
// middleware/auth.go
package middleware

import (
    "strings"
    "my-blog-api/config"
    "my-blog-api/util"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
)

func AuthRequired() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := extractToken(c)
        if token == "" {
            return util.SendError(c, fiber.StatusUnauthorized, "Missing or invalid token", nil)
        }
        
        claims, err := validateToken(token)
        if err != nil {
            return util.SendError(c, fiber.StatusUnauthorized, "Invalid token", err.Error())
        }
        
        // Store user info in context
        c.Locals("user_id", claims["user_id"])
        c.Locals("user_email", claims["email"])
        c.Locals("user_role", claims["role"])
        
        return c.Next()
    }
}

func RoleRequired(roles ...string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userRole := c.Locals("user_role")
        if userRole == nil {
            return util.SendError(c, fiber.StatusUnauthorized, "Authentication required", nil)
        }
        
        roleStr := userRole.(string)
        for _, role := range roles {
            if roleStr == role {
                return c.Next()
            }
        }
        
        return util.SendError(c, fiber.StatusForbidden, "Insufficient permissions", nil)
    }
}
```

### Rate Limiting Middleware

```go
// middleware/rate_limit.go
package middleware

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimit(requests int, duration time.Duration) fiber.Handler {
    return limiter.New(limiter.Config{
        Max:        requests,
        Expiration: duration,
        KeyGenerator: func(c *fiber.Ctx) string {
            return c.IP()
        },
        LimitReached: func(c *fiber.Ctx) error {
            return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
                "error":   true,
                "message": "Rate limit exceeded",
            })
        },
    })
}
```

## Database Operations

### Migrations

Create migration files in the `migrations/` directory:

```sql
-- migrations/001_create_users_table.up.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
```

```sql
-- migrations/001_create_users_table.down.sql
DROP TABLE IF EXISTS users;
```

### Running Migrations

```bash
# Create new migration
oakhouse migrate create create_posts_table

# Run all pending migrations
oakhouse migrate up

# Rollback last migration
oakhouse migrate down

# Check migration status
oakhouse migrate status
```

## Authentication

### JWT Authentication

```go
// service/auth_service.go
package service

import (
    "time"
    "my-blog-api/config"
    "my-blog-api/entity"
    "github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
    GenerateToken(user *entity.User) (string, error)
    ValidateToken(tokenString string) (*jwt.MapClaims, error)
}

type authService struct {
    config *config.Config
}

func NewAuthService(config *config.Config) AuthService {
    return &authService{config: config}
}

func (s *authService) GenerateToken(user *entity.User) (string, error) {
    claims := jwt.MapClaims{
        "user_id":    user.ID,
        "email":      user.Email,
        "role":       user.Role,
        "exp":        time.Now().Add(24 * time.Hour).Unix(),
        "iat":        time.Now().Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.config.JWTSecret))
}
```

### Login Handler

```go
func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var dto auth.LoginDto
    if err := c.BodyParser(&dto); err != nil {
        return util.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
    }
    
    if err := h.validator.Struct(&dto); err != nil {
        return util.SendError(c, fiber.StatusBadRequest, "Validation failed", err.Error())
    }
    
    user, err := h.userService.AuthenticateUser(c.Context(), dto.Email, dto.Password)
    if err != nil {
        return util.SendError(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
    }
    
    token, err := h.authService.GenerateToken(user)
    if err != nil {
        return util.SendError(c, fiber.StatusInternalServerError, "Failed to generate token", err.Error())
    }
    
    return util.SendSuccess(c, "Login successful", fiber.Map{
        "token": token,
        "user":  user,
    })
}
```

## Testing

### Unit Tests

```go
// service/user_service_test.go
package service_test

import (
    "context"
    "testing"
    "my-blog-api/dto/user"
    "my-blog-api/entity"
    "my-blog-api/mocks"
    "my-blog-api/service"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := &mocks.UserRepository{}
    userService := service.NewUserService(mockRepo)
    
    dto := &user.CreateUserDto{
        Name:     "John Doe",
        Email:    "john@example.com",
        Password: "password123",
    }
    
    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil)
    
    // Act
    result, err := userService.CreateUser(context.Background(), dto)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, dto.Name, result.Name)
    assert.Equal(t, dto.Email, result.Email)
    assert.NotEqual(t, dto.Password, result.Password) // Should be hashed
    mockRepo.AssertExpectations(t)
}
```

### Integration Tests

```go
// handler/user_handler_test.go
package handler_test

import (
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "testing"
    "my-blog-api/dto/user"
    "my-blog-api/handler"
    "my-blog-api/mocks"
    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
)

func TestUserHandler_CreateUser(t *testing.T) {
    // Setup
    app := fiber.New()
    mockService := &mocks.UserService{}
    userHandler := handler.NewUserHandler(mockService)
    
    app.Post("/users", userHandler.CreateUser)
    
    // Test data
    dto := user.CreateUserDto{
        Name:     "John Doe",
        Email:    "john@example.com",
        Password: "password123",
    }
    
    jsonData, _ := json.Marshal(dto)
    
    // Mock expectations
    mockService.On("CreateUser", mock.Anything, &dto).Return(&entity.User{
        Name:  dto.Name,
        Email: dto.Email,
    }, nil)
    
    // Execute
    req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    
    resp, err := app.Test(req)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, 200, resp.StatusCode)
    mockService.AssertExpectations(t)
}
```

## Deployment

### Docker Deployment

```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

EXPOSE 8080
CMD ["./main"]
```

```yaml
# docker-compose.yml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=production
      - DB_HOST=postgres
      - REDIS_URL=redis://redis:6379
    depends_on:
      - postgres
      - redis

  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: blog_db
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    volumes:
      - redis_data:/data

volumes:
  postgres_data:
  redis_data:
```

### Kubernetes Deployment

```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: blog-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: blog-api
  template:
    metadata:
      labels:
        app: blog-api
    spec:
      containers:
      - name: blog-api
        image: blog-api:latest
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          value: "postgres-service"
        - name: REDIS_URL
          value: "redis://redis-service:6379"
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
```

## Best Practices

### 1. Error Handling

```go
// Use custom error types
type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Handle errors consistently
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    user, err := h.service.CreateUser(c.Context(), &dto)
    if err != nil {
        switch err.(type) {
        case ValidationError:
            return util.SendError(c, fiber.StatusBadRequest, "Validation failed", err)
        case DuplicateError:
            return util.SendError(c, fiber.StatusConflict, "Resource already exists", err)
        default:
            return util.SendError(c, fiber.StatusInternalServerError, "Internal server error", nil)
        }
    }
    
    return util.SendSuccess(c, "User created successfully", user)
}
```

### 2. Logging

```go
// Use structured logging
import "github.com/sirupsen/logrus"

func (s *userService) CreateUser(ctx context.Context, dto *user.CreateUserDto) (*entity.User, error) {
    logger := logrus.WithFields(logrus.Fields{
        "operation": "CreateUser",
        "email":     dto.Email,
    })
    
    logger.Info("Creating new user")
    
    user, err := s.repo.Create(ctx, &entity.User{...})
    if err != nil {
        logger.WithError(err).Error("Failed to create user")
        return nil, err
    }
    
    logger.WithField("user_id", user.ID).Info("User created successfully")
    return user, nil
}
```

### 3. Validation

```go
// Custom validation tags
func init() {
    validate := validator.New()
    validate.RegisterValidation("password", validatePassword)
}

func validatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    // Check for at least one uppercase, one lowercase, one digit
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
    hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
    
    return hasUpper && hasLower && hasDigit
}

type CreateUserDto struct {
    Password string `json:"password" validate:"required,min=8,password"`
}
```

### 4. Database Transactions

```go
func (s *userService) CreateUserWithProfile(ctx context.Context, dto *user.CreateUserWithProfileDto) error {
    return s.db.DB.Transaction(func(tx *gorm.DB) error {
        // Create user
        user := &entity.User{...}
        if err := tx.Create(user).Error; err != nil {
            return err
        }
        
        // Create profile
        profile := &entity.Profile{
            UserID: user.ID,
            ...
        }
        if err := tx.Create(profile).Error; err != nil {
            return err
        }
        
        return nil
    })
}
```

## Examples

### Complete Blog API Example

This example shows how to build a complete blog API with users, posts, and comments:

```bash
# Create project
oakhouse new blog-api
cd blog-api

# Generate resources
oakhouse generate resource User name:string email:string password:string role:string
oakhouse generate resource Post title:string content:text author_id:uuid published:bool
oakhouse generate resource Comment content:text post_id:uuid author_id:uuid

# Set up relationships in models
# Add to entity/user.go:
# Posts []Post `gorm:"foreignKey:AuthorID"`
# Comments []Comment `gorm:"foreignKey:AuthorID"`

# Add to entity/post.go:
# Author User `gorm:"foreignKey:AuthorID"`
# Comments []Comment `gorm:"foreignKey:PostID"`

# Add to entity/comment.go:
# Post Post `gorm:"foreignKey:PostID"`
# Author User `gorm:"foreignKey:AuthorID"`

# Run migrations
oakhouse migrate up

# Start development server
oakhouse serve
```

### API Endpoints

The generated API will include:

```
# Authentication
POST /api/v1/auth/login
POST /api/v1/auth/register

# Users
GET    /api/v1/users
GET    /api/v1/users/:id
POST   /api/v1/users
PUT    /api/v1/users/:id
DELETE /api/v1/users/:id

# Posts
GET    /api/v1/posts
GET    /api/v1/posts/:id
POST   /api/v1/posts
PUT    /api/v1/posts/:id
DELETE /api/v1/posts/:id

# Comments
GET    /api/v1/comments
GET    /api/v1/comments/:id
POST   /api/v1/comments
PUT    /api/v1/comments/:id
DELETE /api/v1/comments/:id
```

### Sample API Requests

```bash
# Register a new user
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'

# Create a post (with JWT token)
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "My First Post",
    "content": "This is the content of my first post.",
    "published": true
  }'

# Get posts with pagination and filtering
curl "http://localhost:8080/api/v1/posts?page=1&page_size=10&published=true&sort_by=created_at&sort_order=desc"
```

This documentation provides a comprehensive guide to using the Go To Oakhouse framework. For more examples and advanced usage, check out the [examples directory](./examples/) in the repository.
