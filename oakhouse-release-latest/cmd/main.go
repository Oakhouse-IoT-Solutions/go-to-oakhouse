// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"log"

	"oakhouse-release-latest/config"
	"oakhouse-release-latest/adapter"
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

	// Initialize Redis (optional)
	var redisAdapter *adapter.RedisAdapter
	if cfg.RedisURL != "" {
		redisAdapter, err = adapter.NewRedisAdapter(cfg)
		if err != nil {
			log.Printf("Warning: Failed to initialize Redis: %v", err)
		} else {
			log.Println("Redis connected successfully")
		}
	}

	// Create and start server
	server := NewAppServer(cfg, db, redisAdapter)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
