# Go To Oakhouse

A powerful Go framework, designed for rapid API development with clean architecture patterns.

## Features

- 🚀 **Fast Development** - CLI tool for rapid scaffolding
- 🏗️ **Clean Architecture** - Repository, Service, Handler pattern
- 🔧 **Code Generation** - Generate models, handlers, services, repositories
- 🌐 **Fiber Framework** - Built on top of Go Fiber for high performance
- 🗄️ **GORM Integration** - Advanced ORM with scoping support
- ✅ **Auto Validation** - Request validation with struct tags
- 🐳 **Docker Ready** - Production-ready containerization
- 📚 **Comprehensive Documentation** - Detailed guides and examples

## Quick Start

### Installation

```bash
go install github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse@latest
```

### Create New Project

```bash
oakhouse new my-api
cd my-api
```

### Generate Resources

```bash
# Generate a complete resource (model, repository, service, handler, routes)
oakhouse generate resource User name:string email:string age:int

# Generate individual components
oakhouse generate model Product
oakhouse generate handler ProductHandler
oakhouse generate service ProductService
oakhouse generate repository ProductRepository
```

### Start Development Server

```bash
oakhouse serve
```

## Project Structure

```
my-api/
├── cmd/
│   ├── main.go              # Application entry point
│   └── app_server.go        # Server configuration
├── config/
│   └── env_config.go        # Environment configuration
├── adapter/
│   ├── database_adapter.go  # Database connection
│   └── postgres/
│       └── gorm.go         # GORM PostgreSQL adapter
├── handler/                 # HTTP handlers (controllers)
│   ├── user_handler.go
│   └── product_handler.go
├── service/                 # Business logic layer
│   ├── user_service.go
│   └── product_service.go
├── repository/              # Data access layer
│   ├── user_repo.go
│   └── product_repo.go
├── dto/                     # Data Transfer Objects
│   ├── user/
│   │   ├── create_user_dto.go
│   │   ├── update_user_dto.go
│   │   └── get_user_dto.go
│   └── product/
├── scope/                   # GORM scopes for filtering
│   ├── users/
│   └── products/
├── route/                   # Route definitions
│   └── v1.go
├── util/                    # Utility functions
├── .env.example
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

## Configuration

Copy `.env.example` to `.env` and configure your environment:

```env
APP_NAME=my-api
APP_PORT=8080
APP_ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_NAME=my_api_db
DB_USER=postgres
DB_PASSWORD=password

JWT_SECRET=your-secret-key
REDIS_URL=redis://localhost:6379
```

## Usage Examples

### Handler Example

```go
type UserHandler interface {
    Create(ctx *fiber.Ctx) error
    FindAll(ctx *fiber.Ctx) error
    FindById(ctx *fiber.Ctx) error
    Update(ctx *fiber.Ctx) error
    Delete(ctx *fiber.Ctx) error
}

type userHandler struct {
    userService service.UserService
}

func (h *userHandler) Create(ctx *fiber.Ctx) error {
    var req dto.CreateUserDto
    if err := ctx.BodyParser(&req); err != nil {
        return ctx.Status(400).JSON(presenter.ErrorResponse("Invalid request", nil))
    }
    
    user, err := h.userService.Create(ctx.Context(), &req)
    if err != nil {
        return ctx.Status(500).JSON(presenter.ErrorResponse("Failed to create user", nil))
    }
    
    return ctx.Status(201).JSON(presenter.SuccessResponse("User created", user))
}
```

### Service Example

```go
type UserService interface {
    Create(ctx context.Context, req *dto.CreateUserDto) (*entity.User, error)
    FindAll(ctx context.Context, filter *dto.GetUserDto) ([]*entity.User, int64, error)
    FindById(ctx context.Context, id uuid.UUID) (*entity.User, error)
    Update(ctx context.Context, id uuid.UUID, req *dto.UpdateUserDto) error
    Delete(ctx context.Context, id uuid.UUID) error
}

type userService struct {
    repo repository.UserRepository
}

func (s *userService) Create(ctx context.Context, req *dto.CreateUserDto) (*entity.User, error) {
    user := &entity.User{
        Name:  req.Name,
        Email: req.Email,
        Age:   req.Age,
    }
    
    return s.repo.Create(ctx, user)
}
```

### Scope Example

```go
// Filter users by age range
func FilterByAgeRange(minAge, maxAge int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if minAge > 0 {
            db = db.Where("age >= ?", minAge)
        }
        if maxAge > 0 {
            db = db.Where("age <= ?", maxAge)
        }
        return db
    }
}

// Usage in repository
func (r *userRepository) FindAll(ctx context.Context, filter *dto.GetUserDto) ([]*entity.User, int64, error) {
    var users []*entity.User
    var total int64
    
    query := r.db.Model(&entity.User{})
    
    if filter.MinAge != nil && filter.MaxAge != nil {
        query = query.Scopes(FilterByAgeRange(*filter.MinAge, *filter.MaxAge))
    }
    
    query.Count(&total)
    err := query.Offset((filter.Page-1)*filter.PageSize).Limit(filter.PageSize).Find(&users).Error
    
    return users, total, err
}
```

## CLI Commands

### Project Management

```bash
# Create new project
oakhouse new <project-name>

# Start development server with hot reload
oakhouse serve --port 8080

# Build for production
oakhouse build
```

### Code Generation

```bash
# Generate complete resource
oakhouse generate resource <name> [fields...]

# Generate individual components
oakhouse generate model <name>
oakhouse generate handler <name>
oakhouse generate service <name>
oakhouse generate repository <name>
oakhouse generate dto <name>
oakhouse generate scope <name>
oakhouse generate middleware <name>
```

### Database Operations

```bash
# Run migrations
oakhouse migrate up

# Rollback migrations
oakhouse migrate down

# Create new migration
oakhouse migrate create <name>

# Check migration status
oakhouse migrate status
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./handler -v
```

## Deployment

### Docker

```bash
# Build image
docker build -t my-api .

# Run container
docker run -p 8080:8080 my-api
```

### Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## Philosophy

Go To Oakhouse follows these core principles:

### 1. Convention over Configuration
- Sensible defaults for rapid development
- Consistent naming conventions
- Standardized project structure

### 2. Clean Architecture
- Separation of concerns with clear layers
- Dependency injection for testability
- Interface-driven design

### 3. Developer Experience
- Intuitive CLI commands
- Comprehensive code generation
- Hot reload for development

### 4. Performance First
- Built on Go Fiber for speed
- Efficient database queries with GORM
- Optimized for production workloads

### 5. Scalability
- Microservice-ready architecture
- Horizontal scaling support
- Cloud-native deployment

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](https://go-to-oakhouse.dev)
- 💬 [Discord Community](https://discord.gg/go-to-oakhouse)
- 🐛 [Issue Tracker](https://github.com/your-org/go-to-oakhouse/issues)
- 📧 [Email Support](mailto:support@go-to-oakhouse.dev)
