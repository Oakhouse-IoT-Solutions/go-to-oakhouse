# Go To Oakhouse Framework Documentation

## Table of Contents

1. [Installation](#installation)
2. [Quick Start](#quick-start)
3. [Project Structure](#project-structure)
4. [Dependency Injection with Wire](#dependency-injection-with-wire)
5. [CLI Commands](#cli-commands)
6. [Configuration](#configuration)
7. [Models and Entities](#models-and-entities)
8. [Repositories](#repositories)
9. [Services](#services)
10. [Handlers](#handlers)
11. [DTOs (Data Transfer Objects)](#dtos-data-transfer-objects)
12. [Scopes](#scopes)
13. [Middleware](#middleware)
14. [Database Operations](#database-operations)
15. [Authentication](#authentication)
16. [Testing](#testing)
17. [Deployment](#deployment)
18. [Best Practices](#best-practices)
19. [Examples](#examples)

## Installation

### Prerequisites

- Go 1.21 or higher
- PostgreSQL (recommended) or MySQL
- Redis (optional, for caching)

### What's New in v1.19.0

- **Enhanced Architecture**: Improved project structure with comprehensive Wire dependency injection documentation
- **Route Management**: Added detailed route setup and RESTful API design patterns
- **Code Quality**: Enhanced code maintainability with better architectural patterns and examples
- **Documentation Excellence**: Comprehensive documentation updates covering all framework components
- **Developer Experience**: Improved CLI tools and enhanced project scaffolding capabilities

### Install CLI Tool

```bash
go install github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse@v1.19.0
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

# Add database support
oakhouse add database

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

Go To Oakhouse follows a clean architecture pattern with clear separation of concerns. Here's the complete project structure:

```
my-blog-api/
├── cmd/
│   ├── main.go                 # Application entry point
│   ├── app_server.go          # Fiber app server setup
│   ├── wire.go                # Wire dependency injection providers
│   └── wire_gen.go            # Generated Wire code (auto-generated)
├── config/
│   └── env_config.go          # Environment configuration management
├── entity/
│   ├── user.go                # User model/entity
│   └── post.go                # Post model/entity
├── repository/
│   ├── user_repository.go     # User data access layer
│   └── post_repository.go     # Post data access layer
├── service/
│   ├── user_service.go        # User business logic
│   └── post_service.go        # Post business logic
├── handler/
│   ├── user_handler.go        # User HTTP handlers
│   └── post_handler.go        # Post HTTP handlers
├── dto/
│   ├── user/
│   │   ├── create_user_dto.go # User creation request DTO
│   │   ├── update_user_dto.go # User update request DTO
│   │   └── get_user_dto.go    # User response DTO
│   └── post/
│       ├── create_post_dto.go # Post creation request DTO
│       ├── update_post_dto.go # Post update request DTO
│       └── get_post_dto.go    # Post response DTO
├── scope/
│   ├── users/
│   │   └── filter_by_role.go  # User query scopes
│   └── products/
│       ├── filter_by_author.go # Post query scopes
│       └── filter_by_status.go
├── middleware/
│   ├── auth.go                # Authentication middleware
│   └── rate_limit.go          # Rate limiting middleware
├── adapter/
│   ├── database_adapter.go    # Database adapter interface
│   └── postgres/
│       └── gorm.go            # PostgreSQL GORM implementation
├── route/
│   └── v1.go                  # API v1 routes setup
├── util/
│   ├── response.go            # HTTP response utilities
│   └── pagination.go          # Pagination utilities
├── static/
│   └── index.html             # Static files (optional)
├── .env.example               # Environment variables template
├── .env                       # Environment variables (local)
├── docker-compose.yml         # Docker services configuration
├── Dockerfile                 # Application container
├── Makefile                   # Build and development commands
└── go.mod                     # Go module dependencies
```

### Key Architecture Components

#### **1. Clean Architecture Layers**

- **`cmd/`**: Application entry point and server setup
  - `main.go`: Bootstrap application with Wire DI
  - `app_server.go`: Fiber server configuration and middleware setup
  - `wire.go`: Dependency injection provider declarations
  - `wire_gen.go`: Auto-generated Wire dependency injection code

- **`config/`**: Configuration management
  - `env_config.go`: Environment-based configuration loading

- **`entity/`**: Domain models and business entities
  - Core business objects with validation rules
  - Database table representations

- **`repository/`**: Data access layer
  - Database operations and queries
  - Interface-based design for testability
  - GORM integration for ORM operations

- **`service/`**: Business logic layer
  - Core business rules and operations
  - Orchestrates between repositories and handlers
  - Transaction management

- **`handler/`**: HTTP presentation layer
  - REST API endpoints
  - Request/response handling
  - Input validation and error handling

#### **2. Data Transfer Objects (DTOs)**

- **`dto/`**: Request and response data structures
  - `create_*_dto.go`: Input validation for creation
  - `update_*_dto.go`: Input validation for updates
  - `get_*_dto.go`: Response formatting
  - JSON serialization tags and validation rules

#### **3. Query Scopes**

- **`scope/`**: Reusable query filters and conditions
  - Modular query building
  - Chainable query operations
  - Complex filtering logic

#### **4. Infrastructure**

- **`adapter/`**: External service adapters
  - `database_adapter.go`: Database interface abstraction
  - `postgres/gorm.go`: PostgreSQL implementation

- **`middleware/`**: HTTP middleware components
  - Authentication, authorization
  - Rate limiting, CORS, logging

- **`route/`**: API route definitions
  - RESTful endpoint mapping
  - Route grouping and versioning

- **`util/`**: Shared utilities
  - Response formatting helpers
  - Pagination utilities
  - Common helper functions

## Dependency Injection with Wire

### Why We Use Wire for Dependency Injection

Go To Oakhouse leverages [Google Wire](https://github.com/google/wire) for compile-time dependency injection, providing a robust foundation for building maintainable and testable applications. Here's why Wire is the perfect choice for our framework:

#### **1. Compile-Time Safety**

Unlike runtime dependency injection frameworks, Wire generates code at compile time, which means:

- **Zero Runtime Overhead**: No reflection or runtime container lookups
- **Early Error Detection**: Dependency issues are caught during compilation, not at runtime
- **Type Safety**: Full Go type checking ensures your dependencies are correctly wired
- **Performance**: Generated code is as fast as hand-written dependency management

```go
// Wire generates this code for you at compile time
func InitializeApp() (*server.AppServer, func(), error) {
    config := config.Load()
    db, cleanup, err := NewDatabaseAdapter(config)
    if err != nil {
        return nil, nil, err
    }
    appServer := server.NewAppServer(config, db)
    return appServer, cleanup, nil
}
```

#### **2. Explicit Dependency Graph**

Wire makes your application's dependency relationships crystal clear:

```go
// cmd/wire.go - Your dependency providers are explicitly declared
var ProviderSet = wire.NewSet(
    // Core Infrastructure
    config.Load,                    // Configuration provider
    adapter.NewDatabaseAdapter,     // Database connection provider
    
    // Repositories (Data Access Layer)
    repository.NewUserRepository,   // User data access
    repository.NewProductRepository, // Product data access
    
    // Services (Business Logic Layer)
    service.NewUserService,         // User business logic
    service.NewProductService,      // Product business logic
    
    // Handlers (Presentation Layer)
    handler.NewUserHandler,         // User HTTP handlers
    handler.NewProductHandler,      // Product HTTP handlers
    
    // Application Server
    NewAppServer,                   // Fiber app server
)
```

This explicit approach provides several benefits:
- **Transparency**: Anyone can understand the dependency graph by reading the provider set
- **Maintainability**: Easy to modify dependencies without hunting through configuration files
- **Documentation**: The provider set serves as living documentation of your app's architecture

#### **3. Testing Excellence**

Wire's approach makes testing significantly easier:

```go
// Easy to create test-specific dependency graphs
var TestProviderSet = wire.NewSet(
    config.LoadTest,        // Test configuration
    NewMockDatabase,        // Mock database for testing
    server.NewAppServer,    // Real app server
)

// Wire generates: InitializeTestApp() (*server.AppServer, func(), error)
```

**Testing Benefits:**
- **Isolated Testing**: Each component can be tested with mock dependencies
- **Integration Testing**: Easy to wire up real components for integration tests
- **No Test Pollution**: Clean separation between test and production dependencies

#### **4. Scalability and Maintainability**

As your application grows, Wire scales with you:

```go
// Adding new services is straightforward
var ProviderSet = wire.NewSet(
    // Core Infrastructure
    config.Load,
    adapter.NewDatabaseAdapter,
    
    // Repositories (Data Access Layer)
    repository.NewUserRepository,
    repository.NewPostRepository,
    repository.NewEmailRepository,  // New repository added easily
    
    // Services (Business Logic Layer)
    service.NewUserService,
    service.NewPostService,
    service.NewEmailService,        // New service added easily
    
    // Handlers (Presentation Layer)
    handler.NewUserHandler,
    handler.NewPostHandler,
    handler.NewEmailHandler,        // New handler added easily
    
    // Application Server
    NewAppServer,
)
```

**Scalability Benefits:**
- **Modular Growth**: Add new dependencies without refactoring existing code
- **Clear Boundaries**: Each component has well-defined dependencies
- **Refactoring Safety**: Compiler catches breaking changes when restructuring

#### **5. Go-Native Philosophy**

Wire aligns perfectly with Go's design principles:

- **Simplicity**: No magic, no annotations, just plain Go code
- **Explicitness**: Dependencies are clearly declared, not hidden in configuration
- **Performance**: Zero runtime cost, maximum compile-time benefits
- **Tooling**: Full IDE support with autocomplete, refactoring, and navigation

### How Wire Works in Go To Oakhouse

#### **1. Provider Functions**

Each component in your application has a provider function:

```go
// config/config.go
func Load() Config {
    return Config{
        Port:     getEnv("APP_PORT", "8080"),
        Database: getEnv("DB_HOST", "localhost"),
        // ... other config
    }
}

// adapter/database.go
func NewDatabaseAdapter(config config.Config) (*DatabaseAdapter, func(), error) {
    db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
    if err != nil {
        return nil, nil, err
    }
    
    cleanup := func() {
        sqlDB, _ := db.DB()
        sqlDB.Close()
    }
    
    return &DatabaseAdapter{DB: db}, cleanup, nil
}

// server/app_server.go
func NewAppServer(config config.Config, db *adapter.DatabaseAdapter) *AppServer {
    return &AppServer{
        Config: config,
        DB:     db,
        App:    fiber.New(),
    }
}
```

#### **2. Wire Configuration**

The `cmd/wire.go` file defines how components are wired together:

```go
//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
    "your-project/config"
    "your-project/adapter"
    "your-project/server"
)

// ProviderSet is a Wire provider set that includes all dependencies
var ProviderSet = wire.NewSet(
    config.Load,
    NewDatabaseAdapter,
    server.NewAppServer,
)

// InitializeApp initializes the application with all dependencies
func InitializeApp() (*server.AppServer, func(), error) {
    wire.Build(ProviderSet)
    return nil, nil, nil
}
```

#### **3. Code Generation**

When you run `go generate ./cmd` or `make wire-gen`, Wire generates the actual implementation:

```go
// cmd/wire_gen.go (generated by Wire)
func InitializeApp() (*server.AppServer, func(), error) {
    config := config.Load()
    databaseAdapter, cleanup, err := NewDatabaseAdapter(config)
    if err != nil {
        return nil, nil, err
    }
    appServer := server.NewAppServer(config, databaseAdapter)
    return appServer, cleanup, nil
}
```

### Wire vs. Other Dependency Injection Approaches

| Approach | Runtime Cost | Type Safety | Compile-Time Errors | Learning Curve |
|----------|--------------|-------------|--------------------|-----------------|
| **Wire** | ✅ Zero | ✅ Full | ✅ Yes | 🟡 Medium |
| Manual DI | ✅ Zero | ✅ Full | ✅ Yes | ✅ Low |
| Runtime DI | ❌ High | ❌ Limited | ❌ No | ❌ High |
| Service Locator | 🟡 Medium | ❌ Limited | ❌ No | 🟡 Medium |

### Detailed Wire Implementation Guide

#### **Understanding Wire's Dependency Injection Patterns**

Go To Oakhouse implements several key dependency injection patterns using Wire:

##### **1. Provider Pattern**

Provider functions are the building blocks of Wire dependency injection. Each provider is responsible for creating and configuring a specific component:

```go
// Configuration Provider - No dependencies
func Load() *Config {
    return &Config{
        AppPort:    getEnv("APP_PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", ""),
        // ... other config
    }
}

// Database Provider - Depends on Config
func ProvideDatabase(cfg *Config) *adapter.DatabaseAdapter {
    db, err := adapter.NewDatabaseAdapter(cfg)
    if err != nil {
        // Graceful degradation: return nil instead of crashing
        return nil
    }
    return db
}

// Application Server Provider - Depends on Config and Database
func NewAppServer(config *Config, db *adapter.DatabaseAdapter) *AppServer {
    return &AppServer{
        Config: config,
        DB:     db,
        App:    fiber.New(),
    }
}
```

**Provider Pattern Benefits:**
- **Single Responsibility**: Each provider handles one component
- **Error Handling**: Centralized error handling for component initialization
- **Testability**: Easy to mock individual providers
- **Reusability**: Providers can be reused across different configurations

##### **2. Constructor Injection Pattern**

All dependencies are provided through function parameters, making them explicit and testable:

```go
// Handler with explicit service dependencies
func NewUserHandler(userService *service.UserService, authService *service.AuthService) *UserHandler {
    return &UserHandler{
        userService: userService,
        authService: authService,
    }
}

// Service with explicit repository dependencies
func NewUserService(userRepo *repository.UserRepository, emailRepo *repository.EmailRepository) *UserService {
    return &UserService{
        userRepo:  userRepo,
        emailRepo: emailRepo,
    }
}
```

**Constructor Injection Benefits:**
- **Explicit Dependencies**: All dependencies are visible in the function signature
- **Immutable Dependencies**: Dependencies are set once during construction
- **Test-Friendly**: Easy to inject mock dependencies for testing
- **Compile-Time Safety**: Missing dependencies cause compilation errors

##### **3. Layered Architecture with Dependency Flow**

Go To Oakhouse follows a clean layered architecture where dependencies flow in one direction:

```
┌─────────────────┐
│   HTTP Layer    │ ← Handlers (depend on Services)
├─────────────────┤
│  Business Layer │ ← Services (depend on Repositories)
├─────────────────┤
│   Data Layer    │ ← Repositories (depend on Database)
├─────────────────┤
│ Infrastructure  │ ← Database, Config, External APIs
└─────────────────┘
```

**Dependency Flow Rules:**
- **Downward Dependencies**: Each layer only depends on layers below it
- **No Circular Dependencies**: Wire prevents circular dependency issues
- **Interface Segregation**: Each layer exposes only what the layer above needs
- **Dependency Inversion**: Higher layers depend on abstractions, not concretions

#### **Wire Code Generation Process**

##### **1. Wire Configuration (`cmd/wire.go`)**

```go
//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
    "your-project/config"
    "your-project/adapter"
)

// ProviderSet defines the complete dependency graph
// Each provider function is called in the correct order
// based on their dependencies
var ProviderSet = wire.NewSet(
    config.Load,        // Configuration provider (no dependencies)
    ProvideDatabase,    // Database provider (depends on config)
    NewAppServer,       // Application server provider (depends on config + db)
)

// InitializeApp tells Wire what to build
func InitializeApp() (*AppServer, func(), error) {
    wire.Build(ProviderSet)
    return nil, nil, nil  // Wire replaces this implementation
}
```

##### **2. Generated Code (`cmd/wire_gen.go`)**

Wire generates the actual dependency injection code:

```go
// Code generated by Wire. DO NOT EDIT.

func InitializeApp() (*AppServer, func(), error) {
    config := config.Load()                    // Step 1: Load config
    databaseAdapter := ProvideDatabase(config) // Step 2: Create database
    appServer := NewAppServer(config, databaseAdapter) // Step 3: Create server
    
    cleanup := func() {
        // Wire handles cleanup if providers return cleanup functions
    }
    
    return appServer, cleanup, nil
}
```

##### **3. Code Generation Commands**

```bash
# Generate Wire code
go generate ./cmd

# Or use the Makefile
make wire-gen

# Install Wire if not available
go install github.com/google/wire/cmd/wire@latest
```

#### **Advanced Wire Patterns in Go To Oakhouse**

##### **1. Conditional Providers with Graceful Degradation**

```go
// Database provider with graceful degradation
func ProvideDatabase(cfg *Config) *adapter.DatabaseAdapter {
    if cfg.DatabaseURL == "" {
        // Return nil for database-free mode
        log.Println("Running in database-free mode")
        return nil
    }
    
    db, err := adapter.NewDatabaseAdapter(cfg)
    if err != nil {
        log.Printf("Database connection failed: %v", err)
        return nil  // Graceful degradation
    }
    return db
}

// Service that handles nil database gracefully
func NewUserService(db *adapter.DatabaseAdapter) *UserService {
    return &UserService{
        db:     db,
        useDB:  db != nil,  // Flag for database availability
    }
}
```

##### **2. Provider Sets for Modularity**

```go
// Core infrastructure providers
var InfrastructureSet = wire.NewSet(
    config.Load,
    ProvideDatabase,
    ProvideRedis,
)

// Repository layer providers
var RepositorySet = wire.NewSet(
    repository.NewUserRepository,
    repository.NewPostRepository,
    repository.NewEmailRepository,
)

// Service layer providers
var ServiceSet = wire.NewSet(
    service.NewUserService,
    service.NewPostService,
    service.NewEmailService,
)

// Handler layer providers
var HandlerSet = wire.NewSet(
    handler.NewUserHandler,
    handler.NewPostHandler,
    handler.NewAuthHandler,
)

// Complete application provider set
var ProviderSet = wire.NewSet(
    InfrastructureSet,
    RepositorySet,
    ServiceSet,
    HandlerSet,
    NewAppServer,
)
```

##### **3. Testing with Wire**

```go
// test/wire.go - Test-specific provider set
//go:build wireinject
// +build wireinject

// Test provider set with mocks
var TestProviderSet = wire.NewSet(
    config.LoadTest,        // Test configuration
    ProvideMockDatabase,    // Mock database
    NewAppServer,           // Real app server
)

func InitializeTestApp() (*AppServer, func(), error) {
    wire.Build(TestProviderSet)
    return nil, nil, nil
}

// Manual dependency injection for unit tests
func setupTestApp() *AppServer {
    cfg := &config.Config{
        AppPort: "8080",
        // ... test config
    }
    
    mockDB := &MockDatabaseAdapter{}
    
    return NewAppServer(cfg, mockDB)
}
```

### Best Practices with Wire

#### **1. Keep Providers Simple**

```go
// Good: Simple, focused provider
func NewUserService(repo repository.UserRepository) service.UserService {
    return &userService{repo: repo}
}

// Avoid: Complex initialization logic in providers
func NewUserService(config Config) service.UserService {
    // Don't do complex setup here
    // Keep providers focused on dependency injection
}
```

#### **2. Use Interfaces for Flexibility**

```go
// Define interfaces for better testing and flexibility
type UserRepository interface {
    Create(user *entity.User) error
    FindByID(id string) (*entity.User, error)
}

// Provider returns interface, not concrete type
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}
```

#### **3. Group Related Providers**

```go
// Group providers by domain or layer
var DatabaseProviderSet = wire.NewSet(
    NewDatabaseAdapter,
    NewUserRepository,
    NewPostRepository,
)

var ServiceProviderSet = wire.NewSet(
    NewUserService,
    NewPostService,
    NewEmailService,
)

var ProviderSet = wire.NewSet(
    config.Load,
    DatabaseProviderSet,
    ServiceProviderSet,
    server.NewAppServer,
)
```

### Getting Started with Wire

When you create a new Go To Oakhouse project, Wire is automatically configured:

1. **Wire dependency** is added to `go.mod`
2. **Wire configuration** is created in `cmd/wire.go`
3. **Makefile targets** are set up for code generation
4. **Generated code** is automatically created during build

You can regenerate Wire code anytime:

```bash
# Generate Wire code
make wire-gen

# Or manually
go generate ./cmd
```

### Conclusion

Wire represents the perfect balance of power and simplicity for dependency injection in Go. By choosing Wire, Go To Oakhouse provides:

- **Developer Productivity**: Focus on business logic, not dependency management
- **Application Reliability**: Compile-time safety prevents runtime dependency errors
- **Testing Excellence**: Easy mocking and isolated testing
- **Performance**: Zero runtime overhead with maximum compile-time benefits
- **Go Philosophy**: Simple, explicit, and performant dependency management

This foundation enables you to build scalable, maintainable applications with confidence, knowing that your dependency graph is correct, efficient, and easy to understand.

### Example Wire Usage

Here's a practical example of how Wire works in a Go To Oakhouse project:

```go
// cmd/wire.go
//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
    "my-blog-api/config"
    "my-blog-api/adapter"
    "my-blog-api/repository"
    "my-blog-api/service"
    "my-blog-api/handler"
    "my-blog-api/server"
)

// ProviderSet includes all application dependencies
var ProviderSet = wire.NewSet(
    // Core
    config.Load,
    adapter.NewDatabaseAdapter,
    
    // Repositories
    repository.NewUserRepository,
    repository.NewPostRepository,
    
    // Services
    service.NewUserService,
    service.NewPostService,
    
    // Handlers
    handler.NewUserHandler,
    handler.NewPostHandler,
    
    // Server
    server.NewAppServer,
)

// InitializeApp wires up the entire application
func InitializeApp() (*server.AppServer, func(), error) {
    wire.Build(ProviderSet)
    return nil, nil, nil
}
```

```go
// service/user_service.go
type UserService interface {
    CreateUser(dto *dto.CreateUserDTO) (*entity.User, error)
    GetUserByID(id string) (*entity.User, error)
}

type userService struct {
    repo repository.UserRepository
}

// NewUserService is a Wire provider function
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(dto *dto.CreateUserDTO) (*entity.User, error) {
    user := &entity.User{
        Name:  dto.Name,
        Email: dto.Email,
    }
    return s.repo.Create(user)
}
```

```go
// handler/user_handler.go
type UserHandler struct {
    service service.UserService
}

// NewUserHandler is a Wire provider function
func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var dto dto.CreateUserDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }
    
    user, err := h.service.CreateUser(&dto)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    
    return c.JSON(user)
}
```

For more detailed information about Wire's capabilities, advanced usage patterns, and troubleshooting, please refer to the [official Wire documentation](https://github.com/google/wire).

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
# Add database support to your project
oakhouse add database
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

## Routes

### Route Setup

Go To Oakhouse uses a centralized route configuration that brings together all your handlers into a clean API structure:

```go
// route/v1.go
package route

import (
    "my-blog-api/handler"
    "my-blog-api/middleware"
    "github.com/gofiber/fiber/v2"
)

// SetupV1Routes configures all API v1 routes
func SetupV1Routes(app *fiber.App, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {
    // API v1 group
    v1 := app.Group("/api/v1")
    
    // Apply global middleware
    v1.Use(middleware.RateLimit(100, time.Minute)) // 100 requests per minute
    
    // Public routes
    auth := v1.Group("/auth")
    auth.Post("/login", userHandler.Login)
    auth.Post("/register", userHandler.Register)
    
    // Protected routes
    protected := v1.Group("/")
    protected.Use(middleware.AuthRequired())
    
    // User routes
    users := protected.Group("/users")
    users.Get("/", userHandler.ListUsers)                    // GET /api/v1/users
    users.Get("/profile", userHandler.GetProfile)            // GET /api/v1/users/profile
    users.Get("/:id", userHandler.GetUser)                   // GET /api/v1/users/:id
    users.Post("/", userHandler.CreateUser)                  // POST /api/v1/users
    users.Put("/:id", userHandler.UpdateUser)                // PUT /api/v1/users/:id
    users.Delete("/:id", userHandler.DeleteUser)             // DELETE /api/v1/users/:id
    
    // Admin-only user routes
    adminUsers := users.Group("/")
    adminUsers.Use(middleware.RoleRequired("admin"))
    adminUsers.Patch("/:id/status", userHandler.UpdateUserStatus)
    
    // Product routes
    products := protected.Group("/products")
    products.Get("/", productHandler.ListProducts)           // GET /api/v1/products
    products.Get("/:id", productHandler.GetProduct)          // GET /api/v1/products/:id
    products.Post("/", productHandler.CreateProduct)         // POST /api/v1/products
    products.Put("/:id", productHandler.UpdateProduct)       // PUT /api/v1/products/:id
    products.Delete("/:id", productHandler.DeleteProduct)    // DELETE /api/v1/products/:id
}
```

### Route Integration in App Server

```go
// cmd/app_server.go
package main

import (
    "my-blog-api/route"
    "my-blog-api/handler"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

type AppServer struct {
    app         *fiber.App
    userHandler *handler.UserHandler
    productHandler *handler.ProductHandler
}

func NewAppServer(
    userHandler *handler.UserHandler,
    productHandler *handler.ProductHandler,
) *AppServer {
    app := fiber.New(fiber.Config{
        ErrorHandler: customErrorHandler,
    })
    
    // Global middleware
    app.Use(logger.New())
    app.Use(cors.New())
    
    // Setup routes
    route.SetupV1Routes(app, userHandler, productHandler)
    
    return &AppServer{
        app:         app,
        userHandler: userHandler,
        productHandler: productHandler,
    }
}

func (s *AppServer) Start(port string) error {
    return s.app.Listen(":" + port)
}
```

### Route Patterns and Best Practices

#### **1. RESTful API Design**

```go
// Standard REST patterns
GET    /api/v1/users          // List all users
GET    /api/v1/users/:id      // Get specific user
POST   /api/v1/users          // Create new user
PUT    /api/v1/users/:id      // Update entire user
PATCH  /api/v1/users/:id      // Partial user update
DELETE /api/v1/users/:id      // Delete user
```

#### **2. Route Grouping and Middleware**

```go
// Group related routes
api := app.Group("/api")
v1 := api.Group("/v1")
users := v1.Group("/users")

// Apply middleware to groups
protected := v1.Group("/")
protected.Use(middleware.AuthRequired())

admin := protected.Group("/admin")
admin.Use(middleware.RoleRequired("admin"))
```

#### **3. Route Parameters and Validation**

```go
// Handler with parameter validation
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    id := c.Params("id")
    
    // Validate UUID format
    if _, err := uuid.Parse(id); err != nil {
        return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID format", err)
    }
    
    user, err := h.service.GetUser(id)
    if err != nil {
        return util.ErrorResponse(c, fiber.StatusNotFound, "User not found", err)
    }
    
    return util.SuccessResponse(c, "User retrieved successfully", user)
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

# Add database support
oakhouse add database

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
