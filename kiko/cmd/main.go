// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"log"

	"kiko/config"
	"kiko/adapter"
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
