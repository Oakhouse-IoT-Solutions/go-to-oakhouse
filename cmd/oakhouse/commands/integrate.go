// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/utils"
	"github.com/spf13/cobra"
)

// IntegrateCmd creates the 'integrate' command for integrating features to existing projects.
// Provides functionality to integrate database support, authentication, middleware,
// and other features to enhance existing Oakhouse projects.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func IntegrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "integrate",
		Short: "Integrate features to your project",
		Long:  `Integrate various features like authentication, middleware, and more to your existing Oakhouse project.`,
	}

	// Add subcommands
	cmd.AddCommand(integrateRedisCmd())

	return cmd
}

// integrateRedisCmd creates the 'integrate redis' subcommand for adding Redis support
func integrateRedisCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "redis",
		Short: "Integrate Redis caching support",
		Long:  `Add Redis caching support to your Oakhouse project including configuration, connection setup, and caching utilities.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := integrateRedis(); err != nil {
				fmt.Printf("❌ Error integrating Redis: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Redis integration completed successfully!")
		},
	}
}

// integrateRedis adds Redis support to the current project
func integrateRedis() error {
	// Check if we're in an Oakhouse project
	if !isOakhouseProject() {
		return fmt.Errorf("not in an Oakhouse project directory. Please run this command from your project root")
	}

	fmt.Println("🚀 Integrating Redis support...")

	// 1. Update go.mod with Redis dependencies
	if err := addRedisDependencies(); err != nil {
		return fmt.Errorf("failed to add Redis dependencies: %v", err)
	}

	// 2. Update .env.example with Redis configuration
	if err := addRedisEnvConfig(); err != nil {
		return fmt.Errorf("failed to add Redis environment configuration: %v", err)
	}

	// 3. Update config to include Redis fields
	if err := updateConfigForRedis(); err != nil {
		return fmt.Errorf("failed to update config for Redis: %v", err)
	}

	// 4. Create Redis adapter
	if err := createRedisAdapter(); err != nil {
		return fmt.Errorf("failed to create Redis adapter: %v", err)
	}

	// 5. Create Redis utilities
	if err := createRedisUtils(); err != nil {
		return fmt.Errorf("failed to create Redis utilities: %v", err)
	}

	// 6. Update main.go to include Redis initialization
	if err := updateMainGoForRedis(); err != nil {
		return fmt.Errorf("failed to update main.go for Redis: %v", err)
	}

	// 7. Update app_server.go to include Redis adapter
	if err := updateAppServerForRedis(); err != nil {
		return fmt.Errorf("failed to update app_server.go for Redis: %v", err)
	}

	fmt.Println("\n📋 Next steps:")
	fmt.Println("1. Run 'go mod tidy' to download Redis dependencies")
	fmt.Println("2. Update your .env file with Redis configuration")
	fmt.Println("3. Import and use Redis in your handlers and services")

	return nil
}

// isOakhouseProject checks if current directory is an Oakhouse project
func isOakhouseProject() bool {
	// Check for go.mod and typical Oakhouse structure
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return false
	}
	if _, err := os.Stat("cmd"); os.IsNotExist(err) {
		return false
	}
	if _, err := os.Stat("config"); os.IsNotExist(err) {
		return false
	}
	return true
}

// addRedisDependencies adds Redis dependencies to go.mod
func addRedisDependencies() error {
	fmt.Println("📦 Adding Redis dependencies...")

	// Read current go.mod
	goModPath := "go.mod"
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return err
	}

	goModContent := string(content)

	// Check if redis dependency already exists
	if strings.Contains(goModContent, "github.com/redis/go-redis/v9") {
		fmt.Println("✓ Redis dependency already exists")
		return nil
	}

	// Add Redis dependency to require section
	requireIndex := strings.Index(goModContent, "require (")
	if requireIndex == -1 {
		return fmt.Errorf("could not find require section in go.mod")
	}

	// Find the end of require section
	lines := strings.Split(goModContent, "\n")
	var newLines []string
	insideRequire := false
	redisAdded := false

	for _, line := range lines {
		if strings.Contains(line, "require (") {
			insideRequire = true
			newLines = append(newLines, line)
			continue
		}

		if insideRequire && strings.Contains(line, ")") && !redisAdded {
			// Add Redis dependency before closing parenthesis
			newLines = append(newLines, "\tgithub.com/redis/go-redis/v9 v9.3.0")
			redisAdded = true
		}

		newLines = append(newLines, line)

		if insideRequire && strings.Contains(line, ")") {
			insideRequire = false
		}
	}

	// Write updated go.mod
	updatedContent := strings.Join(newLines, "\n")
	return os.WriteFile(goModPath, []byte(updatedContent), 0644)
}

// addRedisEnvConfig adds Redis configuration to .env.example
func addRedisEnvConfig() error {
	fmt.Println("⚙️ Adding Redis environment configuration...")

	envExamplePath := ".env.example"
	content, err := os.ReadFile(envExamplePath)
	if err != nil {
		return err
	}

	envContent := string(content)

	// Check if Redis config already exists
	if strings.Contains(envContent, "REDIS_URL") {
		fmt.Println("✓ Redis configuration already exists")
		return nil
	}

	// Add Redis configuration
	redisConfig := `
# Redis Configuration
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
`

	updatedContent := envContent + redisConfig
	return os.WriteFile(envExamplePath, []byte(updatedContent), 0644)
}

// createRedisAdapter creates Redis adapter file
func createRedisAdapter() error {
	fmt.Println("🔧 Creating Redis adapter...")

	adapterDir := "adapter"
	redisAdapterPath := filepath.Join(adapterDir, "redis_adapter.go")

	// Check if Redis adapter already exists
	if _, err := os.Stat(redisAdapterPath); err == nil {
		fmt.Println("✓ Redis adapter already exists")
		return nil
	}

	// Get project name from go.mod
	projectName, err := getProjectName()
	if err != nil {
		return err
	}

	redisAdapterContent := fmt.Sprintf(templates.RedisAdapterTemplate, projectName)

	return utils.WriteFile(redisAdapterPath, redisAdapterContent, nil)
}

// createRedisUtils creates Redis utility file
func createRedisUtils() error {
	fmt.Println("🔧 Creating Redis utilities...")

	utilDir := "util"
	redisUtilPath := filepath.Join(utilDir, "redis_util.go")

	// Check if Redis util already exists
	if _, err := os.Stat(redisUtilPath); err == nil {
		fmt.Println("✓ Redis utilities already exist")
		return nil
	}

	// Get project name from go.mod
	projectName, err := getProjectName()
	if err != nil {
		return err
	}

	redisUtilContent := fmt.Sprintf(templates.RedisUtilTemplate, projectName)

	return utils.WriteFile(redisUtilPath, redisUtilContent, nil)
}

// updateConfigForRedis updates the config file to include Redis fields if missing
func updateConfigForRedis() error {
	fmt.Println("🔧 Updating config for Redis...")

	configPath := filepath.Join("config", "env_config.go")

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config file not found at %s", configPath)
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	configStr := string(content)

	// Check if Redis fields already exist
	if strings.Contains(configStr, "RedisURL") {
		fmt.Println("✓ Redis configuration already exists in config")
		return nil
	}

	// 1. Patch Config struct
	structInsertIndex := strings.Index(configStr, "type Config struct {")
	if structInsertIndex == -1 {
		return fmt.Errorf("could not find Config struct definition")
	}

	structEndIndex := strings.Index(configStr[structInsertIndex:], "}")
	if structEndIndex == -1 {
		return fmt.Errorf("could not find end of Config struct")
	}

	structBlock := configStr[structInsertIndex : structInsertIndex+structEndIndex]
	if !strings.Contains(structBlock, "RedisURL") {
		redisFields := `
	RedisURL      string
	RedisPassword string
	RedisDB       string`
		configStr = strings.Replace(configStr, structBlock+"}", structBlock+redisFields+"\n}", 1)
	}

	// 2. Patch LoadConfig() return block
	loadFuncIndex := strings.Index(configStr, "func LoadConfig() *Config {")
	if loadFuncIndex == -1 {
		return fmt.Errorf("could not find LoadConfig function")
	}

	returnBlockStart := strings.Index(configStr[loadFuncIndex:], "return &Config{")
	if returnBlockStart == -1 {
		return fmt.Errorf("could not find return &Config block")
	}

	startPos := loadFuncIndex + returnBlockStart
	endPos := strings.Index(configStr[startPos:], "}") + startPos
	if endPos == -1 {
		return fmt.Errorf("could not find end of Config return block")
	}

	// Inject Redis fields
	redisLines := `
		RedisURL:      getEnv("REDIS_URL", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnv("REDIS_DB", "0"),`

	// Insert above the closing `}` of return block
	configStr = configStr[:endPos] + redisLines + configStr[endPos:]

	// Write back
	return utils.WriteFile(configPath, configStr, nil)
}

// updateMainGoForRedis updates the main.go file to include Redis initialization
func updateMainGoForRedis() error {
	mainPath := "cmd/main.go"
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found at %s", mainPath)
	}

	content, err := os.ReadFile(mainPath)
	if err != nil {
		return err
	}

	mainStr := string(content)

	// Check if Redis is already integrated
	if strings.Contains(mainStr, "RedisAdapter") {
		fmt.Println("✓ main.go already contains Redis integration")
		return nil
	}

	// Add Redis initialization after database initialization
	dbInitPattern := `	// Initialize database
	db, err := adapter.InitializeDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}`

	redisInit := `	// Initialize database
	db, err := adapter.InitializeDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize Redis (optional)
	var redisAdapter *adapter.RedisAdapter
	if cfg.RedisURL != "" {
		redisAdapter, err = adapter.NewRedisAdapter(cfg)
		if err != nil {
			log.Fatal("Failed to initialize Redis", err)
		} else {
			log.Println("Redis connected successfully")
		}
	}`

	mainStr = strings.Replace(mainStr, dbInitPattern, redisInit, 1)

	// Update server creation to include Redis adapter
	serverPattern := `server := NewAppServer(cfg, db)`
	redisServerPattern := `server := NewAppServer(cfg, db, redisAdapter)`
	mainStr = strings.Replace(mainStr, serverPattern, redisServerPattern, 1)

	return utils.WriteFile(mainPath, mainStr, nil)
}

// updateAppServerForRedis updates the app_server.go file to include Redis adapter
func updateAppServerForRedis() error {
	appServerPath := "cmd/app_server.go"
	if _, err := os.Stat(appServerPath); os.IsNotExist(err) {
		return fmt.Errorf("app_server.go not found at %s", appServerPath)
	}

	content, err := os.ReadFile(appServerPath)
	if err != nil {
		return err
	}

	appServerStr := string(content)

	// Check if Redis is already integrated
	if strings.Contains(appServerStr, "redisAdapter") {
		fmt.Println("✓ app_server.go already contains Redis integration")
		return nil
	}

	// Update AppServer struct to include Redis adapter
	structPattern := `type AppServer struct {
	app *fiber.App
	cfg *config.Config
	db  *gorm.DB
}`

	redisStruct := `type AppServer struct {
	app          *fiber.App
	cfg          *config.Config
	db           *gorm.DB
	redisAdapter *adapter.RedisAdapter
}`

	appServerStr = strings.Replace(appServerStr, structPattern, redisStruct, 1)

	// Update NewAppServer function signature
	funcPattern := `func NewAppServer(cfg *config.Config, db *gorm.DB) *AppServer {`
	redisFuncPattern := `func NewAppServer(cfg *config.Config, db *gorm.DB, redisAdapter *adapter.RedisAdapter) *AppServer {`
	appServerStr = strings.Replace(appServerStr, funcPattern, redisFuncPattern, 1)

	// Update return statement
	returnPattern := `	return &AppServer{
		app: app,
		cfg: cfg,
		db:  db,
	}`

	redisReturnPattern := `	return &AppServer{
		app:          app,
		cfg:          cfg,
		db:           db,
		redisAdapter: redisAdapter,
	}`

	appServerStr = strings.Replace(appServerStr, returnPattern, redisReturnPattern, 1)

	return utils.WriteFile(appServerPath, appServerStr, nil)
}

// getProjectName extracts project name from go.mod
func getProjectName() (string, error) {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}

	return "", fmt.Errorf("could not find module name in go.mod")
}
