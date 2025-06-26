package main

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
