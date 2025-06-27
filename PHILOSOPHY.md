# Go To Oakhouse Framework Philosophy

## Vision

Go To Oakhouse is a modern, opinionated web framework for Go that brings the developer experience and productivity of frameworks like Laravel and Ruby on Rails to the Go ecosystem. Our mission is to make Go web development faster, more enjoyable, and more accessible without sacrificing performance or type safety.

## Core Principles

### 1. Convention Over Configuration

We believe that sensible defaults and established conventions reduce cognitive load and accelerate development. The framework provides:

- **Standardized Project Structure**: Every Go To Oakhouse project follows the same organizational pattern
- **Naming Conventions**: Predictable naming for models, handlers, services, and repositories
- **Auto-Discovery**: Components are automatically wired together based on naming conventions
- **Sensible Defaults**: Database connections, middleware, and routing work out of the box

```
project/
├── cmd/main.go              # Application entry point
├── config/                  # Configuration management
├── entity/                  # Domain models
├── repository/              # Data access layer
├── service/                 # Business logic layer
├── handler/                 # HTTP handlers
├── dto/                     # Data transfer objects
├── scope/                   # Query scopes
├── middleware/              # Custom middleware
├── adapter/                 # External service adapters
├── route/                   # Route definitions
└── util/                    # Utility functions
```

### 2. Clean Architecture

The framework enforces a clean, layered architecture that promotes:

- **Separation of Concerns**: Each layer has a single responsibility
- **Dependency Inversion**: Higher-level modules don't depend on lower-level modules
- **Testability**: Each layer can be tested in isolation
- **Maintainability**: Changes in one layer don't cascade to others

#### Architecture Layers

1. **Presentation Layer** (Handlers)
   - HTTP request/response handling
   - Input validation
   - Response formatting

2. **Application Layer** (Services)
   - Business logic orchestration
   - Transaction management
   - Cross-cutting concerns

3. **Domain Layer** (Entities)
   - Core business entities
   - Domain rules and invariants
   - Pure business logic

4. **Infrastructure Layer** (Repositories, Adapters)
   - Data persistence
   - External service integration
   - Technical implementations

### 3. Developer Experience First

We prioritize developer productivity and happiness:

#### Powerful CLI Tool
```bash
# Create a new project
oakhouse new my-api

# Generate a complete resource
oakhouse generate resource User name:string email:string age:int

# Generate individual components
oakhouse generate model Product
oakhouse generate handler OrderHandler
oakhouse generate service PaymentService

# Database setup
oakhouse add database

# Development server with hot reload
oakhouse serve
```

#### Intelligent Code Generation
- **Complete CRUD Operations**: Generate models, repositories, services, handlers, and DTOs
- **Type-Safe Code**: All generated code is fully typed
- **Validation Ready**: Built-in validation tags and logic
- **Database Ready**: GORM models with proper relationships

#### Hot Reload Development
- Automatic server restart on code changes
- Fast compilation and startup
- Integrated with popular tools like Air

### 4. Type Safety and Performance

Leveraging Go's strengths:

- **Compile-Time Safety**: Catch errors before runtime
- **Zero-Cost Abstractions**: Framework overhead is minimal
- **Concurrent by Default**: Built on Fiber for high performance
- **Memory Efficient**: Optimized for low memory usage

### 5. Batteries Included

Common functionality works out of the box:

#### Database Integration
- **GORM Integration**: Powerful ORM with database management
- **Multiple Databases**: PostgreSQL, MySQL, SQLite support
- **Connection Pooling**: Optimized database connections
- **Query Scopes**: Reusable query logic

#### Authentication & Authorization
- **JWT Support**: Secure token-based authentication
- **Middleware Integration**: Easy to add auth to routes
- **Role-Based Access**: Built-in permission system

#### Caching & Performance
- **Redis Integration**: Fast caching and sessions
- **Response Caching**: Automatic response caching
- **Rate Limiting**: Built-in rate limiting middleware

#### API Features
- **RESTful by Default**: Standard REST conventions
- **JSON API**: Consistent JSON responses
- **Pagination**: Built-in pagination support
- **Filtering & Sorting**: Query parameter handling

## Design Patterns

### Repository Pattern
Encapsulates data access logic and provides a more object-oriented view of the persistence layer.

```go
type UserRepository interface {
    FindAll(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]entity.User, error)
    FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
    Create(ctx context.Context, user *entity.User) error
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id uuid.UUID) error
}
```

### Service Layer Pattern
Encapsulates business logic and coordinates between different repositories.

```go
type UserService interface {
    CreateUser(ctx context.Context, dto *user.CreateUserDto) (*entity.User, error)
    GetUserProfile(ctx context.Context, userID uuid.UUID) (*entity.User, error)
    UpdateUserProfile(ctx context.Context, userID uuid.UUID, dto *user.UpdateUserDto) error
}
```

### DTO Pattern
Data Transfer Objects for API input/output, separate from domain entities.

```go
type CreateUserDto struct {
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}
```

### Scope Pattern
Reusable query logic that can be composed and applied to different queries.

```go
func FilterByStatus(status string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("status = ?", status)
    }
}

func FilterByDateRange(from, to time.Time) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("created_at BETWEEN ? AND ?", from, to)
    }
}
```

## Configuration Philosophy

### Environment-Based Configuration
All configuration is environment-based with sensible defaults:

```go
type Config struct {
    AppName  string `env:"APP_NAME" default:"go-to-oakhouse"`
    AppPort  string `env:"APP_PORT" default:"8080"`
    DBHost   string `env:"DB_HOST" default:"localhost"`
    DBPort   string `env:"DB_PORT" default:"5432"`
}
```

### Twelve-Factor App Compliance
- Configuration in environment variables
- Explicit dependencies
- Stateless processes
- Port binding
- Graceful shutdown

## Testing Philosophy

### Test Pyramid
1. **Unit Tests**: Test individual components in isolation
2. **Integration Tests**: Test component interactions
3. **End-to-End Tests**: Test complete user workflows

### Testing Tools
- **Testify**: Assertions and mocking
- **Test Containers**: Integration testing with real databases
- **HTTP Testing**: Built-in HTTP test helpers

```go
func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := &mocks.UserRepository{}
    service := NewUserService(mockRepo)
    dto := &user.CreateUserDto{
        Name:  "John Doe",
        Email: "john@example.com",
    }
    
    // Act
    result, err := service.CreateUser(context.Background(), dto)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, dto.Name, result.Name)
}
```

## Deployment Philosophy

### Container-First
Every project includes Docker configuration:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### Cloud-Native
- **Health Checks**: Built-in health endpoints
- **Graceful Shutdown**: Proper signal handling
- **Metrics**: Prometheus-compatible metrics
- **Logging**: Structured logging with levels

## Community and Ecosystem

### Open Source First
- MIT License for maximum flexibility
- Community-driven development
- Extensive documentation and examples
- Active Discord community

### Plugin Architecture
Extensible through plugins:

```go
type Plugin interface {
    Name() string
    Initialize(app *fiber.App) error
    Shutdown() error
}
```

### Middleware Ecosystem
Rich middleware ecosystem for common needs:

- Authentication (JWT, OAuth, Session)
- Rate Limiting
- CORS
- Compression
- Logging
- Metrics
- Caching

## Future Vision

### Roadmap
1. **GraphQL Support**: First-class GraphQL integration
2. **Real-time Features**: WebSocket and Server-Sent Events
3. **Microservices**: Service mesh integration
4. **AI Integration**: Code generation with AI assistance
5. **Performance**: Continuous performance optimization

### Goals
- **10x Developer Productivity**: Make Go web development as fast as Rails
- **Enterprise Ready**: Support for large-scale applications
- **Global Adoption**: Become the go-to framework for Go web development

## Getting Started

Ready to experience the future of Go web development?

```bash
# Install the CLI
go install github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse@v1.16.0

# Create your first project
oakhouse new my-awesome-api
cd my-awesome-api

# Generate your first resource
oakhouse generate resource User name:string email:string

# Start developing
oakhouse serve
```

Welcome to Go To Oakhouse - where Go web development meets developer happiness! 🚀
