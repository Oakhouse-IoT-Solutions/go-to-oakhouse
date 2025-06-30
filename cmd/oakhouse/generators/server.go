package generators

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// startDevServer starts the development server with hot reload support.
// Automatically detects if Air is available for hot reloading, otherwise falls back to standard go run.
// Validates project structure before starting to ensure it's a valid Go To Oakhouse project.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func StartDevServer(host string, port int) error {
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

// buildApplication builds the application for production deployment.
// Compiles the application into a binary for efficient production execution.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func BuildApplication() error {
	fmt.Println("🔨 Building application for production...")
	cmd := exec.Command("go", "build", "-o", "bin/app", "cmd/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// IntegrateDatabase adds database configuration and connection to existing projects.
// Updates environment configuration and modifies app server to require database connection.
// Transforms optional database setup into required database connection for production use.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func IntegrateDatabase() error {
	fmt.Println("📦 Integrating database support...")

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
			"// Initialize database (optional - server can run without it)\n\tvar err error\n\ts.db, err = adapter.NewDatabaseAdapter(s.config)\n\tif err != nil {\n\t\tlog.Printf(\"⚠️  Database connection failed: %v\", err)\n\t\tlog.Println(\"💡 To connect to PostgreSQL, set these environment variables:\")\n\t\tlog.Println(\"   DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME\")\n\t\tlog.Println(\"   Or use: oakhouse integrate database\")\n\t\tlog.Println(\"🚀 Server will continue without database connection\")\n\t\ts.db = nil\n\t}",
			"// Initialize database\n\tvar err error\n\ts.db, err = adapter.NewDatabaseAdapter(s.config)\n\tif err != nil {\n\t\tlog.Fatalf(\"Failed to connect to database: %v\", err)\n\t}")

		if err := os.WriteFile(appServerPath, []byte(updatedContent), 0644); err != nil {
			return fmt.Errorf("failed to update app_server.go: %w", err)
		}
	}

	fmt.Println("✅ Database support configured!")
	return nil
}