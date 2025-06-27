// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"oakhouse-release-v1.16.0/config"
	"oakhouse-release-v1.16.0/adapter"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	// Initialize application server manually for testing
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

// initializeAppManually provides manual dependency injection for testing
func initializeAppManually() (*AppServer, func(), error) {
	// Load configuration
	cfg := config.Load()
	
	// Create database adapter (can be nil for testing)
	var db *adapter.DatabaseAdapter
	var err error
	if cfg != nil {
		db, err = adapter.NewDatabaseAdapter(cfg)
		if err != nil {
			// Continue without database for testing
			log.Printf("Warning: Could not connect to database: %v", err)
			db = nil
		}
	}
	
	// Create app server
	app := NewAppServer(cfg, db)
	
	// Cleanup function
	cleanup := func() {
		if db != nil {
			// Add any cleanup logic here
		}
	}
	
	return app, cleanup, nil
}
