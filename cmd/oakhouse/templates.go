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

const mainGoTemplate = `package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	// Initialize application server
	app := NewAppServer().Init()

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
`

const appServerTemplate = `package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/adapter"
	"{{.ProjectName}}/route"
	"{{.ProjectName}}/middleware"

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

func NewAppServer() *AppServer {
	return &AppServer{}
}

func (s *AppServer) Init() *AppServer {
	// Load configuration
	s.config = config.Load()

	// Initialize database
	var err error
	s.db, err = adapter.NewDatabaseAdapter(s.config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Fiber app
	s.app = fiber.New(fiber.Config{
		AppName: s.config.AppName,
		ErrorHandler: s.errorHandler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Setup middleware
	s.setupMiddleware()

	// Setup routes
	s.setupRoutes()

	return s
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

	// Rate limiting middleware
	s.app.Use(middleware.RateLimit(s.config))
}

func (s *AppServer) setupRoutes() {
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

const envConfigTemplate = `package config

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

const routeTemplate = `package route

import (
	"{{.ProjectName}}/adapter"
	"{{.ProjectName}}/handler"
	"{{.ProjectName}}/repository"
	"{{.ProjectName}}/service"
	"{{.ProjectName}}/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupV1Routes(api fiber.Router, db *adapter.DatabaseAdapter) {
	v1 := api.Group("/v1")

	// Initialize repositories
	// userRepo := repository.NewUserRepository(db)

	// Initialize services
	// userService := service.NewUserService(userRepo)

	// Initialize handlers
	// userHandler := handler.NewUserHandler(userService)

	// Public routes
	public := v1.Group("/")
	// public.Post("/auth/login", authHandler.Login)
	// public.Post("/auth/register", authHandler.Register)

	// Protected routes
	protected := v1.Group("/", middleware.AuthRequired())
	// protected.Get("/users", userHandler.FindAll)
	// protected.Get("/users/:id", userHandler.FindById)
	// protected.Post("/users", userHandler.Create)
	// protected.Put("/users/:id", userHandler.Update)
	// protected.Delete("/users/:id", userHandler.Delete)
}
`

const databaseAdapterTemplate = `package adapter

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

	// Auto-migrate models
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}

	log.Println("✅ Database connected successfully")

	return &DatabaseAdapter{DB: db}, nil
}

func autoMigrate(db *gorm.DB) error {
	// Add your models here for auto-migration
	// return db.AutoMigrate(
	//     &entity.User{},
	//     &entity.Product{},
	// )
	return nil
}

func (d *DatabaseAdapter) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
`

const postgresAdapterTemplate = `package postgres

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

const responseUtilTemplate = `package util

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

const paginationUtilTemplate = `package util

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

const baseScopeTemplate = `package scope

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

const authMiddlewareTemplate = `package middleware

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
	@echo "  migrate-up   Run database migrations"
	@echo "  migrate-down Rollback database migrations"

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

migrate-up:
	@echo "Running migrations..."
	oakhouse migrate up

migrate-down:
	@echo "Rolling back migrations..."
	oakhouse migrate down

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
