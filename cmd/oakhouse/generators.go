package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// createNewProject creates a new Go To Oakhouse project with complete directory structure,
// generates all necessary files from templates, downloads dependencies, and sets up Wire dependency injection.
// It creates a fully functional Go web application with clean architecture patterns.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func createNewProject(projectName string) error {
	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create directory structure
	dirs := []string{
		"cmd",
		"config",
		"adapter",
		"adapter/postgres",
		"handler",
		"service",
		"repository",
		"dto",
		"scope",
		"route",
		"util",
		"middleware",
		"entity",
		"static",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectName, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Generate project files
	files := map[string]string{
		"go.mod":                    goModTemplate,
		".env.example":              envExampleTemplate,
		"Dockerfile":               dockerfileTemplate,
		"docker-compose.yml":       dockerComposeTemplate,
		"cmd/main.go":              mainGoTemplate,
		"cmd/app_server.go":        appServerTemplate,
		"cmd/wire.go":              wireTemplate,

		"config/env_config.go":     envConfigTemplate,
		"route/v1.go":              routeTemplate,
		"adapter/database_adapter.go": databaseAdapterTemplate,
		"adapter/postgres/gorm.go": postgresAdapterTemplate,
		"util/response.go":         responseUtilTemplate,
		"util/pagination.go":       paginationUtilTemplate,
		"scope/base_scope.go":      baseScopeTemplate,
		"middleware/auth.go":       authMiddlewareTemplate,
		"static/index.html":        indexHtmlTemplate,
		"Makefile":                 makefileTemplate,
	}

	for filename, tmpl := range files {
		if err := generateFileFromTemplate(filepath.Join(projectName, filename), tmpl, map[string]string{
			"ProjectName": projectName,
			"ModuleName":  strings.ReplaceAll(projectName, "-", ""),
		}); err != nil {
			return fmt.Errorf("failed to generate %s: %w", filename, err)
		}
	}

	// Download dependencies
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download dependencies: %w", err)
	}

	// Project created successfully
	log.Println("🚀 Project created successfully!")
	log.Printf("📁 Navigate to your project: cd %s", projectName)
	log.Println("🏃 Run your application: go run ./cmd")

	return nil
}

// generateResource generates a complete REST resource including model, repository, service, handler, DTOs and routes.
// This provides a full CRUD implementation following clean architecture principles with proper separation of concerns.
// Returns a list of all created files for user feedback.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateResource(name string, fields []string) ([]string, error) {
	var createdFiles []string
	
	if err := generateModel(name, fields); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("model/%s.go", strings.ToLower(name)))
	
	if err := generateRepository(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("repository/%s_repo.go", strings.ToLower(name)))
	
	if err := generateService(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("service/%s_service.go", strings.ToLower(name)))
	
	if err := generateHandler(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("handler/%s_handler.go", strings.ToLower(name)))
	
	if err := generateDTO(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/create_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/update_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/get_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))
	
	if err := generateRoute(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("route/%s.go", strings.ToLower(name)))
	
	return createdFiles, nil
}

// generateModel generates a GORM model with UUID primary key, timestamps, and soft delete support.
// Creates a struct with proper GORM tags and JSON serialization for database operations.
// Fields are parsed and mapped to appropriate Go types with validation tags.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateModel(name string, fields []string) error {
	filename := fmt.Sprintf("model/%s.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, modelTemplate, map[string]interface{}{
		"ModelName": name,
		"TableName": strings.ToLower(name) + "s",
		"Fields":    parseFields(fields),
	})
}

// generateRepository generates a repository implementation with full CRUD operations.
// Includes context support, GORM scopes, pagination, and proper error handling.
// Follows repository pattern for clean separation between business logic and data access.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateRepository(name string) error {
	filename := fmt.Sprintf("repository/%s_repo.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, repositoryImplTemplate, map[string]interface{}{
		"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

// generateService generates a service layer implementation with business logic operations.
// Handles data transformation between DTOs and models, applies business rules and validation.
// Provides clean interface between handlers and repositories following service pattern.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateService(name string) error {
	filename := fmt.Sprintf("service/%s_service.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, serviceImplTemplate, map[string]interface{}{
		"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

// generateHandler generates a REST API handler with full CRUD endpoints.
// Creates HTTP handlers for Create, Read, Update, Delete operations with proper status codes,
// request validation, error handling, and JSON responses following REST conventions.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateHandler(name string) error {
	filename := fmt.Sprintf("handler/%s_handler.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, handlerTemplate, map[string]interface{}{
		"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

// generateDTO generates Data Transfer Objects for Create, Update, and Get operations.
// Creates separate DTOs with proper validation tags for request/response data transformation.
// Ensures clean separation between API contracts and internal data models.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateDTO(name string) error {
	dtoDir := fmt.Sprintf("dto/%s", strings.ToLower(name))
	if err := os.MkdirAll(dtoDir, 0755); err != nil {
		return fmt.Errorf("failed to create DTO directory: %w", err)
	}

	dtos := map[string]string{
		"create": createDtoTemplate,
		"update": updateDtoTemplate,
		"get":    getDtoTemplate,
	}

	for dtoType, tmpl := range dtos {
		filename := fmt.Sprintf("%s/%s_%s_dto.go", dtoDir, dtoType, strings.ToLower(name))
		if err := generateFileFromTemplate(filename, tmpl, map[string]interface{}{
			"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
			"ModelName":   name,
			"PackageName": strings.ToLower(name),
			"VarName":     strings.ToLower(name),
			"Type":        strings.Title(dtoType),
		}); err != nil {
			return err
		}
	}

	return nil
}

// generateScope generates a GORM scope function for reusable query conditions.
// Creates modular query builders that can be composed and reused across different repository methods.
// Promotes DRY principles and consistent query patterns throughout the application.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateScope(modelName, scopeName string) error {
	scopeDir := fmt.Sprintf("scope/%s", strings.ToLower(modelName))
	if err := os.MkdirAll(scopeDir, 0755); err != nil {
		return fmt.Errorf("failed to create scope directory: %w", err)
	}

	filename := fmt.Sprintf("%s/%s.go", scopeDir, strings.ToLower(scopeName))
	return generateFileFromTemplate(filename, scopeTemplate, map[string]interface{}{
		"ModelName":   modelName,
		"PackageName": strings.ToLower(modelName),
		"ScopeName":   scopeName,
		"FieldName":   scopeName,
	})
}

// generateMiddleware generates HTTP middleware for cross-cutting concerns.
// Creates reusable middleware functions for authentication, logging, CORS, rate limiting,
// and other request/response processing that can be applied to routes or route groups.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateMiddleware(name string) error {
	filename := fmt.Sprintf("middleware/%s.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, middlewareTemplate, map[string]string{
		"Name": name,
	})
}

// getModuleName reads and extracts the module name from the go.mod file.
// Parses the go.mod file to determine the current project's module path,
// which is used for generating proper import statements in generated code.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func getModuleName() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}
	return "", fmt.Errorf("module name not found in go.mod")
}

// generateRoute generates HTTP routes for a resource with full REST endpoints.
// Creates route definitions with proper HTTP methods (GET, POST, PUT, DELETE),
// path parameters, middleware integration, and handler binding for complete API functionality.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateRoute(name string) error {
	// Get project name from go.mod
	projectName, err := getModuleName()
	if err != nil {
		projectName = "your-project" // fallback
	}
	
	// Generate route file
	filename := fmt.Sprintf("route/%s.go", strings.ToLower(name))
	if err := generateFileFromTemplate(filename, resourceRouteTemplate, map[string]interface{}{
		"ProjectName": projectName,
		"Name":        name,
		"LowerName":   strings.ToLower(name),
	}); err != nil {
		return err
	}
	
	// Update v1.go to register the new routes
	return updateV1Routes(name)
}

// updateV1Routes automatically registers new resource routes in the v1 API router.
// Modifies the existing v1.go file to include the new resource routes,
// maintaining proper API versioning and route organization structure.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func updateV1Routes(resourceName string) error {
	v1FilePath := "route/v1.go"
	
	// Check if v1.go exists
	if _, err := os.Stat(v1FilePath); os.IsNotExist(err) {
		return fmt.Errorf("v1.go file not found at %s", v1FilePath)
	}
	
	// Read the current v1.go file
	content, err := os.ReadFile(v1FilePath)
	if err != nil {
		return fmt.Errorf("failed to read v1.go: %v", err)
	}
	
	contentStr := string(content)
	setupCall := fmt.Sprintf("Setup%sRoutes(v1)", resourceName)
	
	// Check if the route is already registered
	if strings.Contains(contentStr, setupCall) {
		return nil // Already registered
	}
	
	// Find the insertion point after the "// Setup resource routes" comment
	commentPattern := "// Setup resource routes"
	commentIndex := strings.Index(contentStr, commentPattern)
	if commentIndex == -1 {
		// If comment doesn't exist, add it before "// Initialize repositories"
		repoComment := "// Initialize repositories"
		repoIndex := strings.Index(contentStr, repoComment)
		if repoIndex == -1 {
			return fmt.Errorf("could not find insertion point in v1.go")
		}
		
		// Insert the comment and route call before the repositories comment
		newContent := contentStr[:repoIndex] + "\t" + commentPattern + "\n\t" + setupCall + "\n\n\t" + contentStr[repoIndex:]
		return os.WriteFile(v1FilePath, []byte(newContent), 0644)
	}
	
	// Find the end of the route setup section
	lines := strings.Split(contentStr, "\n")
	commentLineIndex := -1
	lastRouteLineIndex := -1
	
	// Find the comment line
	for i, line := range lines {
		if strings.Contains(line, commentPattern) {
			commentLineIndex = i
			break
		}
	}
	
	// Find the last route setup line after the comment
	if commentLineIndex != -1 {
		for i := commentLineIndex + 1; i < len(lines); i++ {
			if strings.Contains(lines[i], "Setup") && strings.Contains(lines[i], "Routes(v1)") {
				lastRouteLineIndex = i
			} else if strings.TrimSpace(lines[i]) != "" && !strings.Contains(lines[i], "Setup") {
				// We've reached a non-route line
				break
			}
		}
	}
	
	// Insert the new route call after the last route setup line
	if lastRouteLineIndex != -1 {
		// Insert after the last route line
		lines = append(lines[:lastRouteLineIndex+1], append([]string{"\t" + setupCall}, lines[lastRouteLineIndex+1:]...)...)
	} else {
		// Insert after the comment line
		lines = append(lines[:commentLineIndex+1], append([]string{"\t" + setupCall}, lines[commentLineIndex+1:]...)...)
	}
	
	// Write the updated content back to the file
	updatedContent := strings.Join(lines, "\n")
	return os.WriteFile(v1FilePath, []byte(updatedContent), 0644)
}

// startDevServer starts the development server with hot reload capabilities.
// Automatically detects and uses Air for hot reloading if available, otherwise falls back to standard go run.
// Provides live development experience with automatic server restart on code changes.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func startDevServer(host string, port int) error {
	// Check if we're in a valid Go To Oakhouse project directory
	if _, err := os.Stat("cmd/main.go"); os.IsNotExist(err) {
		return fmt.Errorf("not in a Go To Oakhouse project directory. Please run this command from your project root directory")
	}

	// Check if cmd/app_server.go exists (indicates it's a generated project)
	if _, err := os.Stat("cmd/app_server.go"); os.IsNotExist(err) {
		return fmt.Errorf("this doesn't appear to be a valid Go To Oakhouse project. Missing cmd/app_server.go")
	}

	fmt.Printf("🚀 Starting Go To Oakhouse development server on %s:%d\n", host, port)
	fmt.Println("📁 Watching for file changes...")

	// Check if air is installed for hot reload
	if _, err := exec.LookPath("air"); err == nil {
		cmd := exec.Command("air")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	// Fallback to regular go run
	cmd := exec.Command("go", "run", "./cmd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("APP_PORT=%d", port))
	return cmd.Run()
}



// buildApplication compiles the application for production deployment.
// Creates an optimized binary executable in the bin directory,
// ready for production deployment with all dependencies statically linked.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func buildApplication() error {
	fmt.Println("🔨 Building application for production...")
	cmd := exec.Command("go", "build", "-o", "bin/app", "cmd/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Helper functions

// generateFileFromTemplate creates files from Go templates with dynamic data injection.
// Handles directory creation, template parsing, and file generation with proper error handling.
// Core utility function used by all code generators for consistent file creation.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func generateFileFromTemplate(filename, tmpl string, data interface{}) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Parse template
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	// Execute template
	if err := t.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

// parseFields parses field definitions from command line arguments into structured Field objects.
// Converts string field definitions (name:type format) into Field structs with proper Go types,
// GORM tags, and JSON tags for database and API serialization.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func parseFields(fields []string) []Field {
	var result []Field
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) == 2 {
			fieldName := strings.Title(parts[0])
			fieldType := mapGoType(parts[1])
			lowerName := strings.ToLower(parts[0])
			
			result = append(result, Field{
				Name:    fieldName,
				Type:    fieldType,
				Tag:     fmt.Sprintf(`json:"%s" gorm:"column:%s"`, lowerName, lowerName),
				GormTag: fmt.Sprintf("column:%s", lowerName),
				JsonTag: lowerName,
			})
		}
	}
	return result
}

// mapGoType maps string type names to their corresponding Go type declarations.
// Provides type mapping from user-friendly names to proper Go types,
// supporting common types like string, int, bool, time, uuid with sensible defaults.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func mapGoType(t string) string {
	switch strings.ToLower(t) {
	case "string":
		return "string"
	case "int":
		return "int"
	case "int64":
		return "int64"
	case "float64":
		return "float64"
	case "bool":
		return "bool"
	case "time":
		return "time.Time"
	case "uuid":
		return "uuid.UUID"
	default:
		return "string"
	}
}

// Field represents a struct field
type Field struct {
	Name    string
	Type    string
	Tag     string
	GormTag string
	JsonTag string
}

// addDatabaseSupport adds database configuration and required database connection to existing projects.
// Updates environment configuration and modifies app_server.go to require database connection,
// converting optional database setup to mandatory for projects that need persistent storage.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func addDatabaseSupport() error {
	fmt.Println("📦 Adding database support...")
	
	// Check if we're in a valid project directory
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return fmt.Errorf("not in a Go project directory (go.mod not found)")
	}
	
	// Update .env.example with database variables
	envContent := `# Application
APP_NAME=MyApp
APP_PORT=8080
APP_ENV=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=myapp_db
DB_SSL_MODE=disable
`
	
	if err := os.WriteFile(".env.example", []byte(envContent), 0644); err != nil {
		return fmt.Errorf("failed to update .env.example: %w", err)
	}
	
	// Update app_server.go to require database connection
	appServerPath := "cmd/app_server.go"
	if _, err := os.Stat(appServerPath); err == nil {
		// Read current content
		content, err := os.ReadFile(appServerPath)
		if err != nil {
			return fmt.Errorf("failed to read app_server.go: %w", err)
		}
		
		// Replace optional database logic with required database
		updatedContent := strings.ReplaceAll(string(content), 
			"// Initialize database (optional - server can run without it)\n\tvar err error\n\ts.db, err = adapter.NewDatabaseAdapter(s.config)\n\tif err != nil {\n\t\tlog.Printf(\"⚠️  Database connection failed: %v\", err)\n\t\tlog.Println(\"💡 To connect to PostgreSQL, set these environment variables:\")\n\t\tlog.Println(\"   DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME\")\n\t\tlog.Println(\"   Or use: oakhouse add database\")\n\t\tlog.Println(\"🚀 Server will continue without database connection\")\n\t\ts.db = nil\n\t}",
			"// Initialize database\n\tvar err error\n\ts.db, err = adapter.NewDatabaseAdapter(s.config)\n\tif err != nil {\n\t\tlog.Fatalf(\"Failed to connect to database: %v\", err)\n\t}")
		
		if err := os.WriteFile(appServerPath, []byte(updatedContent), 0644); err != nil {
			return fmt.Errorf("failed to update app_server.go: %w", err)
		}
	}
	
	fmt.Println("✅ Database support configured!")
	return nil
}
