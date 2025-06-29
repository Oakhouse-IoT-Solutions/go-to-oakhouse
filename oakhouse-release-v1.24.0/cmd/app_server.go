// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"fmt"
	"log"

	"oakhouse-release-v1.24.0/config"
	"oakhouse-release-v1.24.0/route"
	"oakhouse-release-v1.24.0/middleware"
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
