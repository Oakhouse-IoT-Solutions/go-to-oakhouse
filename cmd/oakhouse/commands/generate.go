// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"fmt"
	"os"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/generators"
	"github.com/spf13/cobra"
)

// GenerateCmd creates the 'generate' command group for code generation operations.
// Provides subcommands for generating various components like resources, models,
// handlers, services, repositories, DTOs, scopes, middleware, and routes.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate code components",
		Long:  `Generate various code components like models, handlers, services, repositories, DTOs, scopes, middleware, and routes.`,
	}

	// Add all generate subcommands
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

// generateResourceCmd creates the command for generating complete CRUD resources.
// Creates a full set of components including model, handler, service, repository,
// DTOs, routes, and database migration for rapid API development.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateResourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resource [name] [fields...]",
		Short: "Generate a complete CRUD resource (model, handler, service, repository, DTOs, routes)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			resourceName := args[0]
			fields := args[1:] // Get all arguments after the first one as fields
			createdFiles, err := generators.GenerateResource(resourceName, fields)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating resource: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Resource '%s' generated successfully!\n", resourceName)
			fmt.Printf("📁 Created files: %v\n", createdFiles)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}

// generateModelCmd creates the command for generating database models.
// Creates GORM model structs with proper field types, validation tags,
// relationships, and database constraints for clean data modeling.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateModelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model [name] [fields...]",
		Short: "Generate a model",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			modelName := args[0]
			fields := args[1:] // Get all arguments after the first one as fields
			if err := generators.GenerateModel(modelName, fields); err != nil {
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
// Creates REST API handlers with proper HTTP methods, request/response handling,
// validation, error handling, and clean separation of concerns.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateHandlerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "handler [name]",
		Short: "Generate a handler",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			handlerName := args[0]
			if err := generators.GenerateHandler(handlerName); err != nil {
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
			if err := generators.GenerateService(serviceName); err != nil {
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
			if err := generators.GenerateRepository(repoName); err != nil {
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
			if err := generators.GenerateDTO(dtoName); err != nil {
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
			if err := generators.GenerateScope(modelName, scopeName); err != nil {
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
			if err := generators.GenerateMiddleware(middlewareName); err != nil {
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
			if err := generators.GenerateRoute(routeName); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating route: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("🚀 Routes for '%s' generated successfully!\n", routeName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}
	return cmd
}