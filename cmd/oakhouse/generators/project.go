package generators

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/utils"
)

// createNewProject creates a new Go To Oakhouse project with complete directory structure,
// generates all necessary files from templates, downloads dependencies, and sets up Wire dependency injection.
// It creates a fully functional Go web application with clean architecture patterns.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func CreateNewProject(projectName string) error {
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
		"model",
		"static",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectName, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Generate project files
	files := map[string]string{
		"go.mod":             templates.GoModTemplate,
		".env.example":       templates.EnvExampleTemplate,
		"Dockerfile":         templates.DockerfileTemplate,
		"docker-compose.yml": templates.DockerComposeTemplate,
		"cmd/main.go":        templates.MainGoTemplate,
		"cmd/app_server.go":  templates.AppServerTemplate,
		"cmd/wire.go":        templates.WireTemplate,

		"config/env_config.go":        templates.EnvConfigTemplate,
		"route/v1.go":                 templates.RouteTemplate,
		"adapter/database_adapter.go": templates.DatabaseAdapterTemplate,
		"adapter/postgres/gorm.go":    templates.PostgresAdapterTemplate,
		"util/response.go":            templates.ResponseUtilTemplate,
		"util/pagination.go":          templates.PaginationUtilTemplate,
		"scope/base_scope.go":         templates.BaseScopeTemplate,
		"middleware/auth.go":          templates.AuthMiddlewareTemplate,
		"static/index.html":           templates.IndexHtmlTemplate,
		"Makefile":                    templates.MakefileTemplate,
	}

	for filename, tmpl := range files {
		if err := utils.WriteFile(filepath.Join(projectName, filename), tmpl, map[string]string{
			"ProjectName": projectName,
			"ModuleName":  strings.ReplaceAll(projectName, "-", ""),
			"Version":     utils.Version,
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
