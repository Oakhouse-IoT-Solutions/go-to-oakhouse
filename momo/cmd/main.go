// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"log"

	"momo/config"
	"momo/adapter"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database (optional - server can run without it)
	var db, err = adapter.InitializeDatabase(cfg)
	if err != nil {
		log.Printf("⚠️  Database connection failed: %v", err)
		log.Println("💡 To connect to PostgreSQL, set these environment variables:")
		log.Println("   DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME")
		log.Println("   Or use: oakhouse integrate database")
		log.Println("🚀 Server will continue without database connection")
		db = nil
	}

	// Create and start server
	server := NewAppServer(cfg, db)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
