// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"oakhouse-release-v-1.24.0/adapter"
	"oakhouse-release-v-1.24.0/config"

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
