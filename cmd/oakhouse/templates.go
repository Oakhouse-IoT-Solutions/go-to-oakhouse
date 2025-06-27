package main

// Project templates
const goModTemplate = `module {{.ProjectName}}

go 1.21

require (
	github.com/gofiber/fiber/v2 v2.52.0
	github.com/go-playground/validator/v10 v10.16.0
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/google/uuid v1.5.0
	github.com/joho/godotenv v1.4.0
	github.com/redis/go-redis/v9 v9.3.0
	gorm.io/gorm v1.25.5
	gorm.io/driver/postgres v1.5.4
)
`

const envExampleTemplate = `# Application Configuration
APP_NAME={{.ProjectName}}
APP_PORT=8080
APP_ENV=development
APP_DEBUG=true

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME={{.ProjectName}}_db
DB_USER=postgres
DB_PASSWORD=password
DB_SSL_MODE=disable
DB_TIMEZONE=UTC

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRES_IN=24h

# Redis Configuration
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# CORS Configuration
CORS_ALLOWED_ORIGINS=*
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=*

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_DURATION=1m
`

const indexHtmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.ProjectName}} - Go To Oakhouse Project</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            color: #333;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
        }
        
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 2rem;
        }
        
        .hero {
            text-align: center;
            color: white;
            padding: 4rem 0;
        }
        
        .hero h1 {
            font-size: 3.5rem;
            margin-bottom: 1rem;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
        }
        
        .hero p {
            font-size: 1.3rem;
            margin-bottom: 2rem;
            opacity: 0.9;
        }
        
        .version-badge {
            display: inline-block;
            background: rgba(255,255,255,0.2);
            padding: 0.5rem 1rem;
            border-radius: 25px;
            font-weight: bold;
            margin-bottom: 2rem;
        }
        
        .features {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 2rem;
            margin: 3rem 0;
        }
        
        .feature-card {
            background: rgba(255,255,255,0.1);
            padding: 2rem;
            border-radius: 15px;
            backdrop-filter: blur(10px);
            border: 1px solid rgba(255,255,255,0.2);
            color: white;
        }
        
        .feature-card h3 {
            font-size: 1.5rem;
            margin-bottom: 1rem;
            color: #fff;
        }
        
        .feature-card p {
            opacity: 0.9;
        }
        
        .cta-section {
            text-align: center;
            margin: 3rem 0;
        }
        
        .cta-button {
            display: inline-block;
            background: #ff6b6b;
            color: white;
            padding: 1rem 2rem;
            text-decoration: none;
            border-radius: 50px;
            font-weight: bold;
            font-size: 1.1rem;
            transition: transform 0.3s ease;
            box-shadow: 0 4px 15px rgba(255,107,107,0.3);
        }
        
        .cta-button:hover {
            transform: translateY(-2px);
            box-shadow: 0 6px 20px rgba(255,107,107,0.4);
        }
        
        .author-section {
            background: rgba(255,255,255,0.1);
            padding: 2rem;
            border-radius: 15px;
            backdrop-filter: blur(10px);
            border: 1px solid rgba(255,255,255,0.2);
            color: white;
            text-align: center;
            margin: 3rem 0;
        }
        
        .author-section h2 {
            margin-bottom: 1rem;
            color: #fff;
        }
        
        .author-info {
            font-size: 1.1rem;
            opacity: 0.9;
        }
        
        .code-block {
            background: rgba(0,0,0,0.3);
            padding: 1rem;
            border-radius: 8px;
            font-family: 'Courier New', monospace;
            margin: 1rem 0;
            color: #fff;
            overflow-x: auto;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="hero">
            <h1>🏠 {{.ProjectName}}</h1>
            <div class="version-badge">Go To Oakhouse v1.19.0</div>
            <p>Your new Go To Oakhouse project is ready!</p>
        </div>
        
        <div class="features">
            <div class="feature-card">
                <h3>🚀 API Endpoints</h3>
                <p>Access your API at <strong>/api/v1</strong>. Generate new resources with the CLI tool.</p>
            </div>
            
            <div class="feature-card">
                <h3>🏗️ Clean Architecture</h3>
                <p>Repository, Service, Handler pattern with dependency injection for maintainable code.</p>
            </div>
            
            <div class="feature-card">
                <h3>🔧 Code Generation</h3>
                <p>Generate models, handlers, services, repositories, and routes with a single command.</p>
            </div>
            
            <div class="feature-card">
                <h3>🌐 High Performance</h3>
                <p>Built on top of Go Fiber framework for lightning-fast HTTP performance.</p>
            </div>
            
            <div class="feature-card">
                <h3>🗄️ GORM Integration</h3>
                <p>Advanced ORM with scoping support and database management.</p>
            </div>
            
            <div class="feature-card">
                <h3>🎯 Health Check</h3>
                <p>Monitor your application health at <strong>/health</strong> endpoint.</p>
            </div>
        </div>
        
        <div class="cta-section">
            <h2 style="color: white; margin-bottom: 1rem;">Next Steps</h2>
            <div class="code-block">
                # Generate a new resource\n                oakhouse generate resource User name:string email:string\n\n                # Start development server\n                oakhouse serve
            </div>
            <a href="https://github.com/Oakhouse-Technology/go-to-oakhouse" class="cta-button">View Documentation</a>
        </div>
        
        <div class="author-section">
            <h2>👨‍💻 Created with Go To Oakhouse</h2>
            <div class="author-info">
                <strong>Framework by Htet Waiyan</strong><br>
                <em>From Oakhouse Technology</em><br><br>
                <p>Go To Oakhouse makes Go development faster and more enjoyable with clean, maintainable code patterns.</p>
            </div>
        </div>
    </div>
</body>
</html>
`

const dockerfileTemplate = `FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy environment file
COPY --from=builder /app/.env.example .env

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
`

const dockerComposeTemplate = `version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=development
      - DB_HOST=postgres
      - REDIS_URL=redis://redis:6379
    depends_on:
      - postgres
      - redis
    volumes:
      - .:/app
    working_dir: /app

  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: {{.ProjectName}}_db
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  adminer:
    image: adminer
    ports:
      - "8081:8080"
    depends_on:
      - postgres

volumes:
  postgres_data:
  redis_data:
`

const mainGoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{.ProjectName}}/adapter"
	"{{.ProjectName}}/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	// Initialize application server manually
	app, cleanup, err := initializeAppManually()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	defer cleanup()

	// Start server in a goroutine
	go func() {
		if err := app.Run(ctx); err != nil {
			log.Printf("Failed to start server: %v", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	if err := app.Shutdown(ctx); err != nil {
		log.Printf("Failed to shutdown server: %v", err)
		os.Exit(1)
	}

	log.Println("Server stopped")
}

// initializeAppManually provides manual dependency injection for testing and development.
// This function demonstrates how to manually wire dependencies without Wire,
// which is useful for:
// - Testing scenarios where you need precise control over dependencies
// - Development environments where database might not be available
// - Debugging dependency injection issues
// - Understanding the dependency graph without Wire's generated code
//
// Manual vs Wire Dependency Injection:
// - Manual: Explicit control, verbose, error-prone for complex graphs
// - Wire: Automated, concise, compile-time safety, better for production
//
// This function mirrors what Wire generates but allows for customization:
// 1. Load configuration manually
// 2. Optionally create database connection (nil for database-free mode)
// 3. Create application server with dependencies
// 4. Return cleanup function for resource management
//
// Returns:
//   *AppServer - Application server with manually injected dependencies
//   func()     - Cleanup function (empty for database-free setup)
//   error      - Any initialization error (nil in this simple case)
func initializeAppManually() (*AppServer, func(), error) {
	// Step 1: Load configuration
	// This replaces the config.Load provider in Wire's ProviderSet
	cfg := config.Load()

	// Step 2: Database initialization (optional)
	// In this case, we're creating a database-free server for development
	// To add database support, uncomment the following:
	//
	// db, err := adapter.NewDatabaseAdapter(cfg)
	// if err != nil {
	//     return nil, nil, fmt.Errorf("failed to initialize database: %w", err)
	// }
	//
	// For now, we pass nil to demonstrate graceful degradation
	var db *adapter.DatabaseAdapter = nil

	// Step 3: Create application server with dependencies
	// This replaces the NewAppServer provider in Wire's ProviderSet
	// The server will handle the nil database gracefully
	appServer := NewAppServer(cfg, db)

	// Step 4: Define cleanup function
	// This function will be called when the application shuts down
	// It should clean up any resources (database connections, file handles, etc.)
	cleanup := func() {
		// No cleanup needed for database-free setup
		// If database was initialized, you would close it here:
		// if db != nil {
		//     db.Close()
		// }
	}

	return appServer, cleanup, nil
}
`

const appServerTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"context"
	"log"
	"time"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/adapter"
	"{{.ProjectName}}/route"
	// "{{.ProjectName}}/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type AppServer struct {
	app *fiber.App
	config *config.Config
	db *adapter.DatabaseAdapter
}

// NewAppServer is a Wire provider function that creates the main application server.
// This function demonstrates the Constructor Injection pattern where all dependencies
// are provided through the function parameters.
//
// Dependency Injection Benefits:
// - Explicit Dependencies: All required components are clearly visible in the signature
// - Testability: Easy to inject mock dependencies for unit testing
// - Flexibility: Can swap implementations without changing the constructor
// - Immutability: Dependencies are set once during construction
//
// Wire Integration:
// - Wire automatically calls this function with the correct dependencies
// - The config parameter comes from config.Load() provider
// - The db parameter comes from ProvideDatabase() provider
// - Wire ensures proper initialization order based on dependency graph
//
// Parameters:
//   config *config.Config - Application configuration (ports, database settings, etc.)
//   db *adapter.DatabaseAdapter - Database connection wrapper (may be nil if DB unavailable)
//
// Returns:
//   *AppServer - Fully configured application server ready to handle requests
func NewAppServer(config *config.Config, db *adapter.DatabaseAdapter) *AppServer {
	server := &AppServer{
		config: config,
		db: db,
	}

	// Initialize Fiber app
	server.app = fiber.New(fiber.Config{
		AppName: config.AppName,
		ErrorHandler: server.errorHandler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Setup middleware
	server.setupMiddleware()

	// Setup routes
	server.setupRoutes()

	return server
}

func (s *AppServer) setupMiddleware() {
	// Recovery middleware
	s.app.Use(recover.New())

	// Logger middleware
	s.app.Use(logger.New(logger.Config{
		Format: "${time} ${status} - ${method} ${path} ${latency}\n",
	}))

	// Security middleware
	s.app.Use(helmet.New())

	// Compression middleware
	s.app.Use(compress.New())

	// CORS middleware
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: s.config.CorsAllowedOrigins,
		AllowMethods: s.config.CorsAllowedMethods,
		AllowHeaders: s.config.CorsAllowedHeaders,
	}))

	// Rate limiting middleware (uncomment when middleware package is available)
	// s.app.Use(middleware.RateLimit(s.config))
}

func (s *AppServer) setupRoutes() {
	// Serve static files
	s.app.Static("/", "./static")

	// Home page - serve index.html
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./static/index.html")
	})

	// Health check
	s.app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time": time.Now(),
		})
	})

	// API routes
	api := s.app.Group("/api")
	route.SetupV1Routes(api, s.db)
}

func (s *AppServer) errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": true,
		"message": err.Error(),
	})
}

func (s *AppServer) Run(ctx context.Context) error {
	port := s.config.AppPort
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	return s.app.Listen(":" + port)
}

func (s *AppServer) Shutdown(ctx context.Context) error {
	return s.app.Shutdown()
}
`

const envConfigTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package config

import (
	"os"
	"strconv"
)

type Config struct {
	// Application
	AppName string
	AppPort string
	AppEnv  string
	AppDebug bool

	// Database
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
	DBTimezone string

	// JWT
	JWTSecret    string
	JWTExpiresIn string

	// Redis
	RedisURL      string
	RedisPassword string
	RedisDB       int

	// CORS
	CorsAllowedOrigins string
	CorsAllowedMethods string
	CorsAllowedHeaders string

	// Rate Limiting
	RateLimitRequests int
	RateLimitDuration string
}

func Load() *Config {
	return &Config{
		// Application
		AppName:  getEnv("APP_NAME", "go-to-oakhouse"),
		AppPort:  getEnv("APP_PORT", "8080"),
		AppEnv:   getEnv("APP_ENV", "development"),
		AppDebug: getEnvBool("APP_DEBUG", true),

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "oakhouse_db"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
		DBTimezone: getEnv("DB_TIMEZONE", "UTC"),

		// JWT
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),

		// Redis
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		// CORS
		CorsAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
		CorsAllowedMethods: getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS"),
		CorsAllowedHeaders: getEnv("CORS_ALLOWED_HEADERS", "*"),

		// Rate Limiting
		RateLimitRequests: getEnvInt("RATE_LIMIT_REQUESTS", 100),
		RateLimitDuration: getEnv("RATE_LIMIT_DURATION", "1m"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
`

const routeTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package route

import (
	"{{.ProjectName}}/adapter"
	// "{{.ProjectName}}/handler"
	// "{{.ProjectName}}/repository"
	// "{{.ProjectName}}/service"
	// "{{.ProjectName}}/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupV1Routes configures API routes with dependency injection.
// This function demonstrates the Service Locator pattern where dependencies
// are passed down through the routing layer to individual handlers.
//
// Dependency Injection in Routing:
// - Database adapter is injected from the application server
// - Repositories are created with the injected database
// - Services are created with repository dependencies
// - Handlers are created with service dependencies
//
// Layered Architecture Pattern:
// Database -> Repository -> Service -> Handler -> Route
//
// This creates a clean separation of concerns where each layer
// only depends on the layer below it, making the code more
// maintainable and testable.
//
// Parameters:
//   api fiber.Router - Fiber router group for API endpoints
//   db *adapter.DatabaseAdapter - Database connection (may be nil)
func SetupV1Routes(api fiber.Router, db *adapter.DatabaseAdapter) {
	v1 := api.Group("/v1")

	// Health check endpoint - no dependencies required
	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Server is running",
		})
	})

	// Repository Layer - Data Access Objects
	// Repositories encapsulate database operations and provide a clean interface
	// for data access. They depend only on the database adapter.
	//
	// Dependency Injection Pattern: Constructor Injection
	// Each repository receives its dependencies through its constructor
	//
	// Example: Initialize repositories (uncomment when needed)
	// userRepo := repository.NewUserRepository(db)  // Injects database dependency
	// postRepo := repository.NewPostRepository(db)  // Each repo gets same DB instance

	// Service Layer - Business Logic
	// Services contain business logic and orchestrate between repositories.
	// They depend on repositories but are independent of HTTP concerns.
	//
	// Dependency Injection Pattern: Constructor Injection
	// Services receive repository dependencies through their constructors
	//
	// Example: Initialize services (uncomment when needed)
	// userService := service.NewUserService(userRepo)              // Single dependency
	// postService := service.NewPostService(postRepo, userRepo)   // Multiple dependencies
	// emailService := service.NewEmailService()                   // No dependencies

	// Handler Layer - HTTP Request/Response
	// Handlers manage HTTP requests/responses and delegate business logic to services.
	// They depend on services but are independent of data access concerns.
	//
	// Dependency Injection Pattern: Constructor Injection
	// Handlers receive service dependencies through their constructors
	//
	// Example: Initialize handlers (uncomment when needed)
	// userHandler := handler.NewUserHandler(userService)          // Service injection
	// postHandler := handler.NewPostHandler(postService)         // Clean separation
	// authHandler := handler.NewAuthHandler(userService)         // Shared services

	// Public routes (uncomment when needed)
	// public := v1.Group("/")
	// public.Post("/auth/login", authHandler.Login)
	// public.Post("/auth/register", authHandler.Register)

	// Protected routes (uncomment when needed)
	// protected := v1.Group("/", middleware.AuthRequired())
	// protected.Get("/users", userHandler.FindAll)
	// protected.Get("/users/:id", userHandler.FindById)
	// protected.Post("/users", userHandler.Create)
	// protected.Put("/users/:id", userHandler.Update)
	// protected.Delete("/users/:id", userHandler.Delete)
}
`

const databaseAdapterTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package adapter

import (
	"fmt"
	"log"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/adapter/postgres"

	"gorm.io/gorm"
)

type DatabaseAdapter struct {
	DB *gorm.DB
}

func NewDatabaseAdapter(cfg *config.Config) (*DatabaseAdapter, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}



	log.Println("✅ Database connected successfully")

	return &DatabaseAdapter{DB: db}, nil
}



func (d *DatabaseAdapter) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
`

const postgresAdapterTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package postgres

import (
	"fmt"
	"time"

	"{{.ProjectName}}/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
		cfg.DBTimezone,
	)

	var logLevel logger.LogLevel
	if cfg.AppDebug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
`

const responseUtilTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        ` + "`json:\"success\"`" + `
	Message string      ` + "`json:\"message\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
	Error   interface{} ` + "`json:\"error,omitempty\"`" + `
}

type PaginatedResponse struct {
	Success    bool        ` + "`json:\"success\"`" + `
	Message    string      ` + "`json:\"message\"`" + `
	Data       interface{} ` + "`json:\"data\"`" + `
	Pagination Pagination  ` + "`json:\"pagination\"`" + `
}

type Pagination struct {
	Page      int   ` + "`json:\"page\"`" + `
	PageSize  int   ` + "`json:\"page_size\"`" + `
	Total     int64 ` + "`json:\"total\"`" + `
	TotalPage int   ` + "`json:\"total_page\"`" + `
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}

func PaginatedSuccessResponse(message string, data interface{}, pagination Pagination) PaginatedResponse {
	return PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}

func SendSuccess(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(SuccessResponse(message, data))
}

func SendError(c *fiber.Ctx, status int, message string, err interface{}) error {
	return c.Status(status).JSON(ErrorResponse(message, err))
}

func SendPaginatedSuccess(c *fiber.Ctx, message string, data interface{}, pagination Pagination) error {
	return c.JSON(PaginatedSuccessResponse(message, data, pagination))
}
`

const paginationUtilTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PaginationParams struct {
	Page     int
	PageSize int
	Offset   int
}

func GetPaginationParams(c *fiber.Ctx) PaginationParams {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
	}
}

func CalculatePagination(page, pageSize int, total int64) Pagination {
	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))

	return Pagination{
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: totalPage,
	}
}
`

const baseScopeTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package scope

import "gorm.io/gorm"

// Omit excludes specified fields from SELECT
func Omit(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(fields...)
	}
}

// Select includes only specified fields in SELECT
func Select(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields...)
	}
}

// Unscoped includes soft deleted records
func Unscoped() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}
}

// OrderBy adds ORDER BY clause
func OrderBy(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// Limit adds LIMIT clause
func Limit(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

// Offset adds OFFSET clause
func Offset(offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}
`

const authMiddlewareTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package middleware

import (
	"strings"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/util"

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

		return c.Next()
	}
}

func extractToken(c *fiber.Ctx) string {
	auth := c.Get("Authorization")
	if auth == "" {
		return ""
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	cfg := config.Load()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

func RateLimit(cfg *config.Config) fiber.Handler {
	// Implementation would use a rate limiting library
	// For now, just return a pass-through middleware
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
`

const makefileTemplate = `.PHONY: help build run test clean docker-build docker-run

# Variables
APP_NAME={{.ProjectName}}
DOCKER_IMAGE=$(APP_NAME):latest

# Default target
help:
	@echo "Available commands:"
	@echo "  build        Build the application"
	@echo "  run          Run the application"
	@echo "  test         Run tests"
	@echo "  clean        Clean build artifacts"
	@echo "  docker-build Build Docker image"
	@echo "  docker-run   Run Docker container"
	@echo "  dev          Start development server"

build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) cmd/main.go

run: build
	@echo "Running $(APP_NAME)..."
	./bin/$(APP_NAME)

dev:
	@echo "Starting development server..."
	air

test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -cover ./...

clean:
	@echo "Cleaning..."
	rm -rf bin/

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE)

docker-compose-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-compose-down:
	@echo "Stopping services..."
	docker-compose down

install-deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

format:
	@echo "Formatting code..."
	go fmt ./...

lint:
	@echo "Running linter..."
	golangci-lint run
`

const wireTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
// Wire Dependency Injection Configuration
// 
// This file defines the dependency injection setup using Google Wire.
// Wire is a compile-time dependency injection tool that generates code
// to wire your application components together safely and efficiently.
//
// Key Benefits:
// - Compile-time safety: Dependency errors are caught at build time
// - Zero runtime overhead: No reflection or runtime container lookups
// - Type safety: Full Go type checking for all dependencies
// - Explicit dependency graph: Clear visibility of component relationships
//
// How it works:
// 1. Define provider functions for each component
// 2. Group providers in a ProviderSet
// 3. Wire generates the initialization code automatically
// 4. Run 'go generate' or 'make wire-gen' to update generated code

//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"{{.ProjectName}}/config"
	"{{.ProjectName}}/adapter"

	"github.com/google/wire"
)

// ProviderSet defines the complete dependency graph for the application.
// This is the central registry where all component providers are declared.
// 
// Dependency Flow:
// config.Load -> ProvideDatabase -> NewAppServer
//             \                 /
//              \-> (injected) ->/
//
// Each provider function in this set will be called by Wire in the correct
// order to satisfy all dependencies. Wire automatically determines the
// execution order based on the function signatures.
var ProviderSet = wire.NewSet(
	// Configuration Provider
	// Loads environment variables and application settings
	// Returns: *config.Config
	config.Load,
	
	// Database Provider
	// Creates database connection with graceful error handling
	// Depends on: *config.Config
	// Returns: *adapter.DatabaseAdapter
	ProvideDatabase,
	
	// Application Server Provider
	// Creates the main application server with all dependencies
	// Depends on: *config.Config, *adapter.DatabaseAdapter
	// Returns: *AppServer
	NewAppServer,
)

// ProvideDatabase is a Wire provider function that creates a database adapter.
// This function implements the Provider Pattern for dependency injection.
//
// Provider Pattern Benefits:
// - Encapsulates complex initialization logic
// - Handles errors gracefully without crashing the application
// - Allows for different database configurations (with/without DB)
// - Provides a single point of database connection management
//
// Error Handling Strategy:
// - If database connection fails, returns nil instead of crashing
// - Allows the application to start without a database (useful for development)
// - Components should check for nil database before using it
//
// Parameters:
//   cfg *config.Config - Application configuration containing database settings
//
// Returns:
//   *adapter.DatabaseAdapter - Database connection wrapper, or nil if connection fails
func ProvideDatabase(cfg *config.Config) *adapter.DatabaseAdapter {
	// Attempt to create database connection using configuration
	db, err := adapter.NewDatabaseAdapter(cfg)
	if err != nil {
		// Graceful degradation: Return nil database but don't fail application startup
		// This allows the server to run without a database connection
		// Useful for:
		// - Development environments where DB might not be available
		// - Testing scenarios with mock databases
		// - Gradual deployment where DB is added later
		return nil
	}
	return db
}

// InitializeApp is the main Wire injector function that bootstraps the entire application.
// This function signature tells Wire what to build and what dependencies to inject.
//
// Wire Injector Pattern:
// - Function signature defines the desired output (*AppServer)
// - Wire.Build() tells Wire which providers to use (ProviderSet)
// - Wire generates the actual implementation in wire_gen.go
// - The return statement here is just a placeholder - Wire replaces it
//
// Generated Code Flow:
// 1. Wire calls config.Load() to get configuration
// 2. Wire calls ProvideDatabase(config) to get database adapter
// 3. Wire calls NewAppServer(config, database) to create app server
// 4. Wire handles any cleanup functions and error propagation
//
// Returns:
//   *AppServer - Fully initialized application server with all dependencies
//   func()     - Cleanup function to call when shutting down (e.g., close DB connections)
//   error      - Any error that occurred during initialization
func InitializeApp() (*AppServer, func(), error) {
	// Wire replaces this entire function body with generated dependency injection code
	// The actual implementation will be in cmd/wire_gen.go after running 'go generate'
	wire.Build(ProviderSet)
	
	// These return values are placeholders - Wire generates the real implementation
	// The generated code will return properly initialized components
	return &AppServer{}, func() {}, nil
}
`
