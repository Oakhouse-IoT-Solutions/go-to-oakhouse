package generators

import (
	"fmt"
	"os"
	"os/exec"
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
