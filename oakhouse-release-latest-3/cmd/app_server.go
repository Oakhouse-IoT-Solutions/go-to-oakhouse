// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"fmt"
	"log"

	"oakhouse-release-latest-3/adapter"
	"oakhouse-release-latest-3/config"
	"oakhouse-release-latest-3/middleware"
	"oakhouse-release-latest-3/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type AppServer struct {
	app          *fiber.App
	cfg          *config.Config
	db           *gorm.DB
	redisAdapter *adapter.RedisAdapter
}

func NewAppServer(cfg *config.Config, db *gorm.DB, redisAdapter *adapter.RedisAdapter) *AppServer {
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
		app:          app,
		cfg:          cfg,
		db:           db,
		redisAdapter: redisAdapter,
	}
}

func (s *AppServer) Start() error {
	// Setup routes
	route.SetupRoutes(s.app, s.db, s.redisAdapter)

	// Start server
	port := s.cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	return s.app.Listen(fmt.Sprintf(":%s", port))
}
