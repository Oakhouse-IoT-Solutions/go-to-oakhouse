// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
// Wire Dependency Injection Configuration
// 
// This file defines the dependency injection setup using Google Wire.
// Wire is a compile-time dependency injection tool that generates code
// to wire your application components together safely and efficiently.
//
// Key Benefits:
// - Compile-time safety: Dependency errors are caught at build time
// - Zero runtime overhead: No reflection or runtime container lookups
// - Type safety: Full Go type checking for all dependencies
// - Explicit dependency graph: Clear visibility of component relationships
//
// How it works:
// 1. Define provider functions for each component
// 2. Group providers in a ProviderSet
// 3. Wire generates the initialization code automatically
// 4. Run 'go generate' or 'make wire-gen' to update generated code

//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"oakhouse-release-v-1.24.0/config"
	"oakhouse-release-v-1.24.0/adapter"

	"github.com/google/wire"
)

// ProviderSet defines the complete dependency graph for the application.
// This is the central registry where all component providers are declared.
// 
// Dependency Flow:
// config.Load -> ProvideDatabase -> NewAppServer
//             \                 /
//              \-> (injected) ->/
//
// Each provider function in this set will be called by Wire in the correct
// order to satisfy all dependencies. Wire automatically determines the
// execution order based on the function signatures.
var ProviderSet = wire.NewSet(
	// Configuration Provider
	// Loads environment variables and application settings
	// Returns: *config.Config
	config.Load,
	
	// Database Provider
	// Creates database connection with graceful error handling
	// Depends on: *config.Config
	// Returns: *adapter.DatabaseAdapter
	ProvideDatabase,
	
	// Application Server Provider
	// Creates the main application server with all dependencies
	// Depends on: *config.Config, *adapter.DatabaseAdapter
	// Returns: *AppServer
	NewAppServer,
)

// ProvideDatabase is a Wire provider function that creates a database adapter.
// This function implements the Provider Pattern for dependency injection.
//
// Provider Pattern Benefits:
// - Encapsulates complex initialization logic
// - Handles errors gracefully without crashing the application
// - Allows for different database configurations (with/without DB)
// - Provides a single point of database connection management
//
// Error Handling Strategy:
// - If database connection fails, returns nil instead of crashing
// - Allows the application to start without a database (useful for development)
// - Components should check for nil database before using it
//
// Parameters:
//   cfg *config.Config - Application configuration containing database settings
//
// Returns:
//   *adapter.DatabaseAdapter - Database connection wrapper, or nil if connection fails
func ProvideDatabase(cfg *config.Config) *adapter.DatabaseAdapter {
	// Attempt to create database connection using configuration
	db, err := adapter.NewDatabaseAdapter(cfg)
	if err != nil {
		// Graceful degradation: Return nil database but don't fail application startup
		// This allows the server to run without a database connection
		// Useful for:
		// - Development environments where DB might not be available
		// - Testing scenarios with mock databases
		// - Gradual deployment where DB is added later
		return nil
	}
	return db
}

// InitializeApp is the main Wire injector function that bootstraps the entire application.
// This function signature tells Wire what to build and what dependencies to inject.
//
// Wire Injector Pattern:
// - Function signature defines the desired output (*AppServer)
// - Wire.Build() tells Wire which providers to use (ProviderSet)
// - Wire generates the actual implementation in wire_gen.go
// - The return statement here is just a placeholder - Wire replaces it
//
// Generated Code Flow:
// 1. Wire calls config.Load() to get configuration
// 2. Wire calls ProvideDatabase(config) to get database adapter
// 3. Wire calls NewAppServer(config, database) to create app server
// 4. Wire handles any cleanup functions and error propagation
//
// Returns:
//   *AppServer - Fully initialized application server with all dependencies
//   func()     - Cleanup function to call when shutting down (e.g., close DB connections)
//   error      - Any error that occurred during initialization
func InitializeApp() (*AppServer, func(), error) {
	// Wire replaces this entire function body with generated dependency injection code
	// The actual implementation will be in cmd/wire_gen.go after running 'go generate'
	wire.Build(ProviderSet)
	
	// These return values are placeholders - Wire generates the real implementation
	// The generated code will return properly initialized components
	return &AppServer{}, func() {}, nil
}
