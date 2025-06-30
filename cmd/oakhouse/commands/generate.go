// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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

// Helper functions for enhanced resource generation

// validateResourceName checks if the resource name follows Go naming conventions
func validateResourceName(name string) error {
	if name == "" {
		return fmt.Errorf("resource name cannot be empty")
	}

	// Check if name starts with uppercase letter and contains only alphanumeric characters
	matched, _ := regexp.MatchString(`^[A-Z][a-zA-Z0-9]*$`, name)
	if !matched {
		return fmt.Errorf("resource name must start with uppercase letter and contain only alphanumeric characters")
	}

	// Check for reserved Go keywords
	reservedWords := []string{"break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select", "struct", "switch", "type", "var"}
	for _, word := range reservedWords {
		if strings.ToLower(name) == word {
			return fmt.Errorf("resource name cannot be a Go reserved keyword: %s", word)
		}
	}

	return nil
}

// validateFields checks if field specifications are valid
func validateFields(fields []string) error {
	validTypes := map[string]bool{
		"string": true, "int": true, "int32": true, "int64": true,
		"uint": true, "uint32": true, "uint64": true,
		"float32": true, "float64": true, "float": true,
		"bool": true, "time.Time": true, "text": true,
		"[]string": true, "[]int": true, "[]float64": true,
	}

	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid field format '%s', expected 'name:type'", field)
		}

		fieldName := strings.TrimSpace(parts[0])
		fieldType := strings.TrimSpace(parts[1])

		if fieldName == "" {
			return fmt.Errorf("field name cannot be empty in '%s'", field)
		}

		// Check field name format
		matched, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9]*$`, fieldName)
		if !matched {
			return fmt.Errorf("field name '%s' must start with letter and contain only alphanumeric characters", fieldName)
		}

		if !validTypes[fieldType] {
			return fmt.Errorf("unsupported field type '%s' in field '%s'", fieldType, field)
		}
	}

	return nil
}

// runInteractiveMode prompts user for resource details
func runInteractiveMode(resourceName string, fields []string) (string, []string, error) {
	reader := bufio.NewReader(os.Stdin)

	// Get resource name if not provided
	if resourceName == "" {
		fmt.Print("Enter resource name (e.g., User, Product): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", nil, err
		}
		resourceName = strings.TrimSpace(input)
	}

	fmt.Printf("\nConfiguring resource: %s\n", resourceName)
	fmt.Println("Enter fields (format: name:type). Press Enter with empty line to finish.")
	fmt.Println("Supported types: string, int, int32, int64, uint, uint32, uint64, float32, float64, float, bool, time.Time, text, []string, []int, []float64")
	fmt.Println("")

	var interactiveFields []string
	for {
		fmt.Printf("Field %d: ", len(interactiveFields)+1)
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", nil, err
		}

		field := strings.TrimSpace(input)
		if field == "" {
			break
		}

		// Validate field format
		if err := validateFields([]string{field}); err != nil {
			fmt.Printf("❌ %v. Please try again.\n", err)
			continue
		}

		interactiveFields = append(interactiveFields, field)
		fmt.Printf("✅ Added field: %s\n", field)
	}

	// Combine existing fields with interactive fields
	allFields := append(fields, interactiveFields...)

	fmt.Printf("\n📋 Resource Summary:\n")
	fmt.Printf("   Name: %s\n", resourceName)
	fmt.Printf("   Fields: %v\n", allFields)
	fmt.Print("\nProceed with generation? (y/N): ")

	confirm, err := reader.ReadString('\n')
	if err != nil {
		return "", nil, err
	}

	if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
		return "", nil, fmt.Errorf("generation cancelled by user")
	}

	return resourceName, allFields, nil
}

// showDryRunPreview shows what files would be generated
func showDryRunPreview(resourceName string, fields []string) {
	fmt.Printf("Resource: %s\n", resourceName)
	fmt.Printf("Fields: %v\n\n", fields)

	fmt.Println("Files that would be generated:")
	files := []string{
		fmt.Sprintf("model/%s.go", strings.ToLower(resourceName)),
		fmt.Sprintf("adapter/repository/%s_repository.go", strings.ToLower(resourceName)),
		fmt.Sprintf("adapter/repository/%s_repository_impl.go", strings.ToLower(resourceName)),
		fmt.Sprintf("adapter/service/%s_service.go", strings.ToLower(resourceName)),
		fmt.Sprintf("adapter/service/%s_service_impl.go", strings.ToLower(resourceName)),
		fmt.Sprintf("adapter/handler/%s_handler.go", strings.ToLower(resourceName)),
		fmt.Sprintf("adapter/dto/%s_dto.go", strings.ToLower(resourceName)),
		fmt.Sprintf("adapter/route/%s_route.go", strings.ToLower(resourceName)),
	}

	for i, file := range files {
		fmt.Printf("   %d. %s\n", i+1, file)
	}

	fmt.Printf("\nTotal: %d files\n", len(files))
}

// checkForConflicts checks if files already exist
func checkForConflicts(resourceName string) []string {
	var conflicts []string
	lowerName := strings.ToLower(resourceName)

	filesToCheck := []string{
		fmt.Sprintf("model/%s.go", lowerName),
		fmt.Sprintf("adapter/repository/%s_repository.go", lowerName),
		fmt.Sprintf("adapter/repository/%s_repository_impl.go", lowerName),
		fmt.Sprintf("adapter/service/%s_service.go", lowerName),
		fmt.Sprintf("adapter/service/%s_service_impl.go", lowerName),
		fmt.Sprintf("adapter/handler/%s_handler.go", lowerName),
		fmt.Sprintf("adapter/dto/%s_dto.go", lowerName),
		fmt.Sprintf("adapter/route/%s_route.go", lowerName),
	}

	for _, file := range filesToCheck {
		if _, err := os.Stat(file); err == nil {
			absPath, _ := filepath.Abs(file)
			conflicts = append(conflicts, absPath)
		}
	}

	return conflicts
}

// generateResourceCmd creates the command for generating complete CRUD resources.
// Creates a full set of components including model, handler, service, repository,
// DTOs, routes, and database migration for rapid API development.
// Enhanced with interactive mode, validation, dry run, and progress feedback.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateResourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resource [name] [fields...]",
		Short: "Generate a complete CRUD resource (model, handler, service, repository, DTOs, routes)",
		Long: `Generate a complete CRUD resource with all necessary components.

This command creates:
- Model with GORM annotations
- Repository with CRUD operations
- Service layer with business logic
- HTTP handlers with REST endpoints
- DTOs for data transfer
- Routes configuration

Examples:
  oakhouse generate resource User name:string email:string age:int
  oakhouse generate resource Product title:string price:float description:text
  oakhouse generate resource --interactive
  oakhouse generate resource --dry-run User name:string`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Get flags
			interactive, _ := cmd.Flags().GetBool("interactive")
			dryRun, _ := cmd.Flags().GetBool("dry-run")
			verbose, _ := cmd.Flags().GetBool("verbose")
			skipValidation, _ := cmd.Flags().GetBool("skip-validation")
			force, _ := cmd.Flags().GetBool("force")

			resourceName := args[0]
			fields := args[1:]

			// Input validation
			if !skipValidation {
				if err := validateResourceName(resourceName); err != nil {
					fmt.Fprintf(os.Stderr, "❌ Invalid resource name: %v\n", err)
					os.Exit(1)
				}
				if err := validateFields(fields); err != nil {
					fmt.Fprintf(os.Stderr, "❌ Invalid field specification: %v\n", err)
					os.Exit(1)
				}
			}

			// Interactive mode
			if interactive {
				var err error
				resourceName, fields, err = runInteractiveMode(resourceName, fields)
				if err != nil {
					fmt.Fprintf(os.Stderr, "❌ Interactive mode failed: %v\n", err)
					os.Exit(1)
				}
			}

			// Dry run mode
			if dryRun {
				fmt.Printf("🔍 Dry run mode - showing what would be generated:\n\n")
				showDryRunPreview(resourceName, fields)
				return
			}

			// Check for existing files
			if !force {
				if conflicts := checkForConflicts(resourceName); len(conflicts) > 0 {
					fmt.Printf("⚠️  The following files already exist:\n")
					for _, file := range conflicts {
						fmt.Printf("   - %s\n", file)
					}
					fmt.Printf("\nUse --force to overwrite existing files.\n")
					os.Exit(1)
				}
			}

			// Progress feedback
			if verbose {
				fmt.Printf("🚀 Starting resource generation for '%s'...\n", resourceName)
				fmt.Printf("📋 Fields: %v\n", fields)
			}

			// Generate resource
			createdFiles, err := generators.GenerateResource(resourceName, fields)
			if err != nil {
				fmt.Fprintf(os.Stderr, "❌ Error generating resource '%s': %v\n", resourceName, err)
				fmt.Fprintf(os.Stderr, "\n💡 Troubleshooting tips:\n")
				fmt.Fprintf(os.Stderr, "   - Ensure you're in a valid Go project directory\n")
				fmt.Fprintf(os.Stderr, "   - Check that field syntax is correct (name:type)\n")
				fmt.Fprintf(os.Stderr, "   - Verify write permissions in the target directory\n")
				os.Exit(1)
			}

			// Success output
			fmt.Printf("\n✅ Resource '%s' generated successfully!\n", resourceName)
			fmt.Printf("📁 Created %d files:\n", len(createdFiles))
			for i, file := range createdFiles {
				fmt.Printf("   %d. %s\n", i+1, file)
			}

			// Next steps
			fmt.Printf("\n🎯 Next steps:\n")
			fmt.Printf("   1. Review the generated files\n")
			fmt.Printf("   2. Run database migrations if needed\n")
			fmt.Printf("   3. Update your main.go to register the routes\n")
			fmt.Printf("   4. Test your new API endpoints\n")
			fmt.Printf("\n🏡 Proudly Created by Htet Waiyan From Oakhouse\n")
		},
	}

	// Add flags
	cmd.Flags().BoolP("interactive", "i", false, "Run in interactive mode to specify fields step by step")
	cmd.Flags().Bool("dry-run", false, "Show what would be generated without creating files")
	cmd.Flags().BoolP("verbose", "v", false, "Enable verbose output with progress information")
	cmd.Flags().Bool("skip-validation", false, "Skip input validation (use with caution)")
	cmd.Flags().BoolP("force", "f", false, "Overwrite existing files without confirmation")

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

// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateHandlerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "handler [name]",
		Short: "Generate a handler",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			handlerName := args[0]
			if err := generators.GenerateSimpleHandler(handlerName); err != nil {
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
			if err := generators.GenerateService(serviceName, []string{}); err != nil {
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
			if err := generators.GenerateDTO(dtoName, []string{}); err != nil {
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
