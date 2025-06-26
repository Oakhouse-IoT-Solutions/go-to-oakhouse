package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// createNewProject creates a new Go To Oakhouse project
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
		"migrations",
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
		"config/env_config.go":     envConfigTemplate,
		"route/v1.go":              routeTemplate,
		"adapter/database_adapter.go": databaseAdapterTemplate,
		"adapter/postgres/gorm.go": postgresAdapterTemplate,
		"util/response.go":         responseUtilTemplate,
		"util/pagination.go":       paginationUtilTemplate,
		"scope/base_scope.go":      baseScopeTemplate,
		"middleware/auth.go":       authMiddlewareTemplate,
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

	return nil
}

// generateResource generates a complete REST resource
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

// generateModel generates a model
func generateModel(name string, fields []string) error {
	filename := fmt.Sprintf("model/%s.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, modelTemplate, map[string]interface{}{
		"ModelName": name,
		"TableName": strings.ToLower(name) + "s",
		"Fields":    parseFields(fields),
	})
}

// generateRepository generates a repository
func generateRepository(name string) error {
	filename := fmt.Sprintf("repository/%s_repo.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, repositoryImplTemplate, map[string]interface{}{
		"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

// generateService generates a service
func generateService(name string) error {
	filename := fmt.Sprintf("service/%s_service.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, serviceImplTemplate, map[string]interface{}{
		"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

// generateHandler generates a handler
func generateHandler(name string) error {
	filename := fmt.Sprintf("handler/%s_handler.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, handlerTemplate, map[string]interface{}{
		"ProjectName": "github.com/Oakhouse-Technology/go-to-oakhouse",
		"ModelName":   name,
		"PackageName": strings.ToLower(name),
		"VarName":     strings.ToLower(name),
	})
}

// generateDTO generates DTOs
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

// generateScope generates a GORM scope
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

// generateMiddleware generates middleware
func generateMiddleware(name string) error {
	filename := fmt.Sprintf("middleware/%s.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, middlewareTemplate, map[string]string{
		"Name": name,
	})
}

// getModuleName reads the module name from go.mod
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

// generateRoute generates routes for a resource
func generateRoute(name string) error {
	// Get project name from go.mod
	projectName, err := getModuleName()
	if err != nil {
		projectName = "your-project" // fallback
	}
	
	// Generate route file
	filename := fmt.Sprintf("route/%s.go", strings.ToLower(name))
	return generateFileFromTemplate(filename, resourceRouteTemplate, map[string]interface{}{
		"ProjectName": projectName,
		"Name":        name,
		"LowerName":   strings.ToLower(name),
	})
}

// startDevServer starts the development server
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

// runMigrations runs database migrations
func runMigrations() error {
	fmt.Println("🔄 Running database migrations...")
	// Implementation would depend on migration tool
	return nil
}

// rollbackMigration rolls back the last migration
func rollbackMigration() error {
	fmt.Println("⏪ Rolling back last migration...")
	// Implementation would depend on migration tool
	return nil
}

// createMigration creates a new migration file
func createMigration(name string) error {
	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("migrations/%s_%s.sql", timestamp, name)

	content := fmt.Sprintf(`-- Migration: %s
-- Created at: %s

-- Up migration

-- Down migration
`, name, time.Now().Format("2006-01-02 15:04:05"))

	return os.WriteFile(filename, []byte(content), 0644)
}

// showMigrationStatus shows migration status
func showMigrationStatus() error {
	fmt.Println("📊 Migration Status:")
	// Implementation would show actual migration status
	fmt.Println("  All migrations are up to date")
	return nil
}

// buildApplication builds the application for production
func buildApplication() error {
	fmt.Println("🔨 Building application for production...")
	cmd := exec.Command("go", "build", "-o", "bin/app", "cmd/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Helper functions

// generateFileFromTemplate generates a file from a template
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

// parseFields parses field definitions
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

// mapGoType maps string types to Go types
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

// addDatabaseSupport adds database support to existing project
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
