// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Project setup templates
const GoModTemplate = `module {{.ProjectName}}

go 1.21

require (
	github.com/gofiber/fiber/v2 v2.52.0
	github.com/google/uuid v1.5.0
	github.com/joho/godotenv v1.4.0
	gorm.io/driver/postgres v1.5.4
	gorm.io/gorm v1.25.5
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
`

const EnvExampleTemplate = `# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME={{.ProjectName}}_db
DB_SSL_MODE=disable

# Server Configuration
PORT=8080
ENV=development

# JWT Configuration
JWT_SECRET=your-secret-key-here
JWT_EXPIRES_IN=24h

# CORS Configuration
CORS_ALLOWED_ORIGINS=*
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=*
`

const DockerfileTemplate = `FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

EXPOSE 8080

CMD ["./main"]
`

const DockerComposeTemplate = `version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=password
      - DB_NAME={{.ProjectName}}_db
      - DB_SSL_MODE=disable
    depends_on:
      - postgres
    networks:
      - {{.ProjectName}}_network

  postgres:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB={{.ProjectName}}_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - {{.ProjectName}}_network

volumes:
  postgres_data:

networks:
  {{.ProjectName}}_network:
    driver: bridge
`

const MainGoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"log"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/adapter"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := adapter.InitializeDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create and start server
	server := NewAppServer(cfg, db)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
`

const AppServerTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"fmt"
	"log"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/route"
	"{{.ProjectName}}/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type AppServer struct {
	app *fiber.App
	cfg *config.Config
	db  *gorm.DB
}

func NewAppServer(cfg *config.Config, db *gorm.DB) *AppServer {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	// Custom middleware
	app.Use(middleware.AuthMiddleware())

	return &AppServer{
		app: app,
		cfg: cfg,
		db:  db,
	}
}

func (s *AppServer) Start() error {
	// Setup routes
	route.SetupRoutes(s.app, s.db)

	// Start server
	port := s.cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	return s.app.Listen(fmt.Sprintf(":%s", port))
}
`

const WireTemplate = `//go:build wireinject
// +build wireinject

// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"{{.ProjectName}}/config"
	"{{.ProjectName}}/adapter"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Wire providers
var AppSet = wire.NewSet(
	config.LoadConfig,
	adapter.InitializeDatabase,
	NewAppServer,
)

// InitializeApp initializes the application with dependency injection
func InitializeApp() (*AppServer, error) {
	wire.Build(AppSet)
	return &AppServer{}, nil
}
`

const EnvConfigTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package config

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	Port       string
	Env        string
	JWTSecret  string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "{{.ProjectName}}_db"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
		Port:       getEnv("PORT", "8080"),
		Env:        getEnv("ENV", "development"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
`

const RouteTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "{{.ProjectName}} API is running",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Add your resource routes here
	// Example:
	// SetupUserRoutes(api, db)
	// SetupProductRoutes(api, db)
}
`

const DatabaseAdapterTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package adapter

import (
	"fmt"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/adapter/postgres"
	"gorm.io/gorm"
)

// InitializeDatabase initializes the database connection
func InitializeDatabase(cfg *config.Config) (*gorm.DB, error) {
	return postgres.NewGormDB(cfg)
}

// GetDSN returns the database connection string
func GetDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)
}
`

const PostgresAdapterTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package postgres

import (
	"fmt"
	"log"

	"{{.ProjectName}}/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewGormDB creates a new GORM database connection
func NewGormDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return db, nil
}
`

const ResponseUtilTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// SuccessResponse returns a standardized success response
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(fiber.Map{
		"requestId": uuid.New(),
		"success":   true,
		"message":   message,
		"data":      data,
	})
}

// ErrorResponse returns a standardized error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err error) error {
	response := fiber.Map{
		"requestId": uuid.New(),
		"success":   false,
		"message":   message,
	}

	if err != nil {
		response["error"] = err.Error()
	}

	return c.Status(statusCode).JSON(response)
}

// PaginatedResponse returns a standardized paginated response
func PaginatedResponse(c *fiber.Ctx, message string, data interface{}, pagination interface{}) error {
	return c.JSON(fiber.Map{
		"requestId":  uuid.New(),
		"success":    true,
		"message":    message,
		"data":       data,
		"pagination": pagination,
	})
}
`

const PaginationUtilTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package util

import "math"

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	CurrentPage int   ` + "`json:\"current_page\"`" + `
	PageSize    int   ` + "`json:\"page_size\"`" + `
	TotalPages  int   ` + "`json:\"total_pages\"`" + `
	TotalItems  int64 ` + "`json:\"total_items\"`" + `
	HasNext     bool  ` + "`json:\"has_next\"`" + `
	HasPrev     bool  ` + "`json:\"has_prev\"`" + `
}

// CalculatePagination calculates pagination metadata
func CalculatePagination(page, pageSize int, totalItems int64) PaginationMeta {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	return PaginationMeta{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
	}
}

// GetOffset calculates the database offset for pagination
func GetOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}
`

const BaseScopeTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package scope

import (
	"time"

	"gorm.io/gorm"
)

// DateRangeScope applies date range filtering with pointer types for consistency
func DateRangeScope(startDate, endDate *time.Time, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startDate != nil {
			db = db.Where(column+" >= ?", *startDate)
		}
		if endDate != nil {
			db = db.Where(column+" <= ?", *endDate)
		}
		return db
	}
}

// SearchScope applies text search filtering
func SearchScope(search, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			return db.Where(column+" ILIKE ?", "%"+search+"%")
		}
		return db
	}
}

// StatusScope applies status filtering
func StatusScope(status, column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where(column+" = ?", status)
		}
		return db
	}
}
`

const AuthMiddlewareTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware handles authentication
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip auth for health check and public routes
		if c.Path() == "/health" || c.Path() == "/" {
			return c.Next()
		}

		// Add your authentication logic here
		// For now, just pass through
		return c.Next()
	}
}
`

const IndexHtmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.ProjectName}} API</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            margin: 0;
            padding: 40px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .container {
            text-align: center;
            max-width: 600px;
        }
        h1 {
            font-size: 3rem;
            margin-bottom: 1rem;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
        }
        p {
            font-size: 1.2rem;
            margin-bottom: 2rem;
            opacity: 0.9;
        }
        .api-link {
            display: inline-block;
            padding: 12px 24px;
            background: rgba(255,255,255,0.2);
            border: 2px solid rgba(255,255,255,0.3);
            border-radius: 8px;
            color: white;
            text-decoration: none;
            font-weight: 600;
            transition: all 0.3s ease;
        }
        .api-link:hover {
            background: rgba(255,255,255,0.3);
            transform: translateY(-2px);
        }
        .footer {
            margin-top: 3rem;
            opacity: 0.7;
            font-size: 0.9rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 {{.ProjectName}}</h1>
        <p>Your API is up and running!</p>
        <a href="/api/v1" class="api-link">Explore API</a>
        <div class="footer">
            <p>🏡 Proudly Created with Oakhouse</p>
        </div>
    </div>
</body>
</html>
`

const MakefileTemplate = `.PHONY: build run test clean docker-build docker-run

# Build the application
build:
	go build -o bin/{{.ProjectName}} cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -t {{.ProjectName}} .

# Run with Docker Compose
docker-run:
	docker-compose up --build

# Run with Docker Compose in background
docker-up:
	docker-compose up -d --build

# Stop Docker Compose
docker-down:
	docker-compose down

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate code
generate:
	go generate ./...

# Database migration (add your migration commands here)
migrate-up:
	# Add your migration up commands

migrate-down:
	# Add your migration down commands
`