package main

import (
	"context"
	"log"
	"time"

	"oakhouse-release-v1.11.0/config"
	"oakhouse-release-v1.11.0/adapter"
	"oakhouse-release-v1.11.0/route"
	// "oakhouse-release-v1.11.0/middleware"

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

	// Initialize database (optional - server can run without it)
	var err error
	s.db, err = adapter.NewDatabaseAdapter(s.config)
	if err != nil {
		log.Printf("⚠️  Database connection failed: %v", err)
		log.Println("💡 To connect to PostgreSQL, set these environment variables:")
		log.Println("   DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME")
		log.Println("   Or use: oakhouse add database")
		log.Println("🚀 Server will continue without database connection")
		s.db = nil
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
