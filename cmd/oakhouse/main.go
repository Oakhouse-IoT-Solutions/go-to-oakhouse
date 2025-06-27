package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "1.9.0"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "oakhouse",
		Short:   "Go To Oakhouse - A powerful Go framework for rapid API development",
		Long:    `Go To Oakhouse is a Go framework for building APIs with clean architecture patterns.

Created by Htet Waiyan From Oakhouse`,
		Version: version,
	}

	// Add subcommands
	rootCmd.AddCommand(newCmd())
	rootCmd.AddCommand(generateCmd())
	rootCmd.AddCommand(addCmd())
	rootCmd.AddCommand(serveCmd())
	rootCmd.AddCommand(dbSetupCmd())
	rootCmd.AddCommand(buildCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// newCmd creates a new project
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
			fmt.Printf("✅ Project '%s' created successfully!\n", projectName)
			fmt.Printf("\nNext steps:\n")
			fmt.Printf("  cd %s\n", projectName)
			fmt.Printf("  cp .env.example .env\n")
			fmt.Printf("  oakhouse serve\n")
		},
	}
	return cmd
}

// generateCmd handles code generation
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

// addCmd handles adding features to existing projects
func addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add features to existing project",
	}

	// Add subcommands
	cmd.AddCommand(addDatabaseCmd())

	return cmd
}

// addDatabaseCmd adds database support to existing project
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
			fmt.Println("   2. Run 'oakhouse db-setup' to initialize database")
		},
	}
	return cmd
}

// generateResourceCmd generates a complete resource
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
			fmt.Printf("✅ Resource '%s' generated successfully!\n", resourceName)
			fmt.Println("\n📁 Created files:")
			for _, file := range createdFiles {
				fmt.Printf("   %s\n", file)
			}
		},
	}
	return cmd
}

// generateModelCmd generates a model
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
			fmt.Printf("✅ Model '%s' generated successfully!\n", modelName)
		},
	}
	return cmd
}

// generateHandlerCmd generates a handler
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
			fmt.Printf("✅ Handler '%s' generated successfully!\n", handlerName)
		},
	}
	return cmd
}

// generateServiceCmd generates a service
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
			fmt.Printf("✅ Service '%s' generated successfully!\n", serviceName)
		},
	}
	return cmd
}

// generateRepositoryCmd generates a repository
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
			fmt.Printf("✅ Repository '%s' generated successfully!\n", repoName)
		},
	}
	return cmd
}

// generateDTOCmd generates DTOs
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
			fmt.Printf("✅ DTOs for '%s' generated successfully!\n", dtoName)
		},
	}
	return cmd
}

// generateScopeCmd generates a scope
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
			fmt.Printf("✅ Scope '%s' for '%s' generated successfully!\n", scopeName, modelName)
		},
	}
	return cmd
}

// generateMiddlewareCmd generates middleware
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
			fmt.Printf("✅ Middleware '%s' generated successfully!\n", middlewareName)
		},
	}
	return cmd
}

// generateRouteCmd generates routes for a resource
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
			fmt.Printf("✅ Routes for '%s' generated successfully!\n", routeName)
		},
	}
	return cmd
}

// serveCmd starts the development server
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

// dbSetupCmd sets up the database
func dbSetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db-setup",
		Short: "Initialize database connection and setup",
		Run: func(cmd *cobra.Command, args []string) {
			if err := setupDatabase(); err != nil {
				fmt.Fprintf(os.Stderr, "Error setting up database: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Database setup completed successfully!")
		},
	}
	return cmd
}

// buildCmd builds the application for production
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
