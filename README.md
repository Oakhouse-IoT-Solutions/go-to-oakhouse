# 🏠  The Oakhouse Way to Build Faster with Go
![The Oakhouse Way](assets/the_oakhouse_way.png)


## Welcome to Go To Oakhouse

**Richard Feynman** 🏆, a Nobel Prize–winning physicist, was known not just for his genius—but for how clearly he explained complex ideas.
He believed that if you can't explain something simply, you don't really understand it.
His method, the **Feynman Technique** 🎯, was all about breaking things down, identifying gaps, and refining your explanation until it truly made sense.
That philosophy is at the heart of **Oakhouse** 🏠.

**Go to Oakhouse** means more than just using a framework.
It means going back to that place— 🌟
Where curiosity led the way.
Where you weren't afraid to ask dumb questions. ❓
Where you stayed up late, learned with friends, and built things just for the fun of it. 🌙👥
The name comes from that house— 🏡
Where we lived after university.
We were broke, hungry, and foolish— 💸🍕😅
But full of ideas. 💡
**Oakhouse** brings that spirit to Go.

I hate complexity for the sake of complexity. 😤
And if you're new to Go, or coming from another language, you've probably felt that pain.
**Oakhouse** 🏠 is a framework built to make your start with Go easier.
It's for beginners and developers coming from other languages. 👨‍💻👩‍💻
It helps you skip repetitive work, avoid boilerplate code, and build real projects faster—without reinventing the wheel. ⚡
More than that,
**Oakhouse** helps you conquer complexity. ⚔️
Just like **Julius Caesar** said:
*"I came, I saw, I conquered."* 🏛️
Here, you'll come, you'll learn, and you'll conquer Go. 🎯
Like **Laravel** for PHP or **Rails** for Ruby,
**Oakhouse** is that—but for Go.
Simple. Curious. Honest. ✨

---

A powerful Go framework, designed for rapid API development with clean architecture patterns.

## Features

- 🚀 **Fast Development** - CLI tool for rapid scaffolding
- 🏗️ **Clean Architecture** - Repository, Service, Handler pattern
- 🔧 **Code Generation** - Generate models, handlers, services, repositories
- 🌐 **Fiber Framework** - Built on top of Go Fiber for high performance
- 🗄️ **GORM Integration** - Advanced ORM with scoping support
- 🔴 **Redis Integration** - Built-in Redis caching with intelligent cache management
- ✅ **Auto Validation** - Request validation with struct tags
- 🎯 **Simplified Handlers** - Generate lightweight handlers with text responses for rapid prototyping
- 🐳 **Docker Ready** - Production-ready containerization
- 📚 **Comprehensive Documentation** - Detailed guides and examples

## What's New in v1.30.0

- 🔴 **Redis Integration** - Comprehensive Redis caching support with intelligent cache management
- ⚡ **Performance Optimization** - Built-in caching strategies for FindAll and FindById operations
- 🔄 **Cache Invalidation** - Automatic cache invalidation for Create, Update, and Delete operations
- ⚙️ **Redis Configuration** - Easy Redis setup with environment variables
- 📚 **Enhanced Documentation** - Complete Redis integration guides and examples

## Quick Start

### Installation

```bash
go install github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse@v1.30.0
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
```

### Start Development Server

```bash
# Server starts with postgres database by default
oakhouse serve
```

### Database Support

By default, new projects are generated **with postgres database support**.

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

```bash

# Generate a handler
oakhouse generate handler UserHandler

```

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

These simplified handlers allow you to:
- ✅ Start testing API endpoints immediately
- ✅ Prototype without database setup
- ✅ Focus on route structure and API design
- ✅ Gradually add full implementation later

### Service Example

```bash

# Generate a service
oakhouse generate handler UserService

```

```go
type UserService interface {
    Create(ctx context.Context, req *dto.CreateUserDto) (*model.User, error)
    FindAll(ctx context.Context, filter *dto.GetUserDto) ([]*model.User, int64, error)
    FindById(ctx context.Context, id uuid.UUID) (*model.User, error)
    Update(ctx context.Context, id uuid.UUID, req *dto.UpdateUserDto) error
    Delete(ctx context.Context, id uuid.UUID) error
}

type userService struct {
    repo repository.UserRepository
}

func (s *userService) Create(ctx context.Context, req *dto.CreateUserDto) (*model.User, error) {
user := &model.User{
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
func (r *userRepository) FindAll(ctx context.Context, filter *dto.GetUserDto) ([]*model.User, int64, error) {
    var users []*model.User
    var total int64
    
    query := r.db.Model(&model.User{})
    
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

# Start development server 
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
oakhouse generate dto <name>
oakhouse generate scope <name>
oakhouse generate middleware <name>
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
