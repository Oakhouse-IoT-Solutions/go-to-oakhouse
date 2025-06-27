package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "1.21.0"
)

// main initializes and executes the Oakhouse CLI application with all available commands.
// Sets up the root command with version information and registers all subcommands
// for project creation, code generation, feature addition, serving, and building.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func main() {
	rootCmd := &cobra.Command{
		Use:     "oakhouse",
		Short:   "🚀 Go To Oakhouse - A powerful Go framework for rapid API development",
		Long:    `🚀 Go To Oakhouse is a Go framework for building APIs with clean architecture patterns.

🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡

Features:
• Clean Architecture patterns
• Rapid API development
• Auto-generated CRUD operations
• Built-in authentication middleware
• Database support
• Docker containerization`,
		Version: version,
	}

	// Add subcommands
	rootCmd.AddCommand(newCmd())
	rootCmd.AddCommand(generateCmd())
	rootCmd.AddCommand(addCmd())
	rootCmd.AddCommand(serveCmd())
	rootCmd.AddCommand(buildCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// newCmd creates the 'new' command for generating new Oakhouse projects.
// Initializes a complete project structure with all necessary files, dependencies,
// and configuration for rapid API development with clean architecture patterns.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func newCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new [project-name]",
		Short: "Create a new Go To Oakhouse project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			if err := createNewProject(projectName); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating project: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("\n🚀 Project '%s' created successfully!\n", projectName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n\n")
			fmt.Printf("Next steps:\n")
			fmt.Printf("  cd %s\n", projectName)
			fmt.Printf("  cp .env.example .env\n")
			fmt.Printf("  oakhouse serve\n\n")
			fmt.Printf("🚀 Happy coding with Oakhouse! 🏡\n")
		},
	}
	return cmd
}

// generateCmd creates the 'generate' command group for code generation operations.
// Provides subcommands for generating models, handlers, services, repositories, DTOs,
// scopes, middleware, routes, and complete resources with proper CLI organization.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "Generate code components",
		Aliases: []string{"gen", "g"},
	}

	// Add generate subcommands
	cmd.AddCommand(generateResourceCmd())
	cmd.AddCommand(generateModelCmd())
	cmd.AddCommand(generateHandlerCmd())
	cmd.AddCommand(generateServiceCmd())
	cmd.AddCommand(generateRepositoryCmd())
	cmd.AddCommand(generateDTOCmd())
	cmd.AddCommand(generateScopeCmd())
	cmd.AddCommand(generateMiddlewareCmd())
	cmd.AddCommand(generateRouteCmd())

	return cmd
}

// addCmd creates the 'add' command group for adding features to existing projects.
// Allows developers to enhance existing projects with additional capabilities
// like database support, authentication, and other framework features.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add features to existing project",
	}

	// Add subcommands
	cmd.AddCommand(addDatabaseCmd())

	return cmd
}

// addDatabaseCmd creates the command for adding database support to existing projects.
// Configures database connection requirements, updates environment variables,
// and modifies application configuration for persistent data storage.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func addDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Add database support to existing project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := addDatabaseSupport(); err != nil {
				fmt.Fprintf(os.Stderr, "Error adding database support: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Database support added successfully!")
			fmt.Println("💡 Don't forget to:")
			fmt.Println("   1. Set database environment variables in .env")
		},
	}
	return cmd
}

// generateResourceCmd creates the command for generating complete REST resources.
// Generates model, repository, service, handler, DTOs, and routes in one command,
// providing a full CRUD implementation following clean architecture principles.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateResourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resource [name] [fields...]",
		Short: "Generate a complete resource (model, repository, service, handler, routes)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			resourceName := args[0]
			fields := args[1:]
			createdFiles, err := generateResource(resourceName, fields)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating resource: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Resource '%s' generated successfully!\n", resourceName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n\n")
			fmt.Println("📁 Created files:")
			for _, file := range createdFiles {
				fmt.Printf("   %s\n", file)
			}
		},
	}
	return cmd
}

// generateModelCmd creates the command for generating GORM models.
// Creates database models with UUID primary keys, timestamps, soft delete support,
// and proper GORM tags for database operations and JSON serialization.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateModelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model [name] [fields...]",
		Short: "Generate a model",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			modelName := args[0]
			fields := args[1:]
			if err := generateModel(modelName, fields); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating model: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Model '%s' generated successfully!\n", modelName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateHandlerCmd creates the command for generating HTTP handlers.
// Creates REST API handlers with full CRUD endpoints, proper status codes,
// request validation, error handling, and JSON responses following REST conventions.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateHandlerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "handler [name]",
		Short: "Generate a handler",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			handlerName := args[0]
			if err := generateHandler(handlerName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating handler: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Handler '%s' generated successfully!\n", handlerName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateServiceCmd creates the command for generating service layer implementations.
// Creates business logic services with data transformation, validation, and clean interfaces
// between handlers and repositories following the service pattern.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service [name]",
		Short: "Generate a service",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			serviceName := args[0]
			if err := generateService(serviceName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating service: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Service '%s' generated successfully!\n", serviceName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateRepositoryCmd creates the command for generating repository implementations.
// Creates data access layer with full CRUD operations, context support, GORM scopes,
// pagination, and proper error handling following the repository pattern.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateRepositoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "repository [name]",
		Short:   "Generate a repository",
		Aliases: []string{"repo"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoName := args[0]
			if err := generateRepository(repoName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating repository: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Repository '%s' generated successfully!\n", repoName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateDTOCmd creates the command for generating Data Transfer Objects.
// Creates separate DTOs for Create, Update, and Get operations with proper validation tags
// for clean separation between API contracts and internal data models.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateDTOCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dto [name]",
		Short: "Generate DTOs (Create, Update, Get)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dtoName := args[0]
			if err := generateDTO(dtoName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating DTO: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 DTOs for '%s' generated successfully!\n", dtoName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateScopeCmd creates the command for generating GORM scopes.
// Creates reusable query condition functions that can be composed and reused
// across different repository methods, promoting DRY principles.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateScopeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scope [name] [scope-name]",
		Short: "Generate a GORM scope",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			modelName := args[0]
			scopeName := args[1]
			if err := generateScope(modelName, scopeName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating scope: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Scope '%s' for '%s' generated successfully!\n", scopeName, modelName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateMiddlewareCmd creates the command for generating HTTP middleware.
// Creates reusable middleware functions for cross-cutting concerns like authentication,
// logging, CORS, rate limiting, and other request/response processing.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateMiddlewareCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "middleware [name]",
		Short: "Generate middleware",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			middlewareName := args[0]
			if err := generateMiddleware(middlewareName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating middleware: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Middleware '%s' generated successfully!\n", middlewareName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateRouteCmd creates the command for generating HTTP routes.
// Creates route definitions with proper HTTP methods, path parameters,
// middleware integration, and handler binding for complete API functionality.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateRouteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "route [name]",
		Short: "Generate routes for a resource",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			routeName := args[0]
			if err := generateRoute(routeName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating route: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Routes for '%s' generated successfully!\n", routeName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// serveCmd creates the command for starting the development server.
// Provides hot reload capabilities using Air if available, with configurable host and port,
// enabling live development experience with automatic server restart on code changes.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func serveCmd() *cobra.Command {
	var port int
	var host string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the development server with hot reload",
		Run: func(cmd *cobra.Command, args []string) {
			if err := startDevServer(host, port); err != nil {
				fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	cmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to bind the server to")

	return cmd
}



// buildCmd creates the command for building the application for production.
// Compiles the application into an optimized binary executable ready for deployment,
// with all dependencies statically linked for production environments.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func buildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the application for production",
		Run: func(cmd *cobra.Command, args []string) {
			if err := buildApplication(); err != nil {
				fmt.Fprintf(os.Stderr, "Error building application: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Application built successfully!")
		},
	}
	return cmd
}
