package generators

import (
	"fmt"
	"os"
	"strings"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/templates"
	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/utils"
)

// generateRoute generates REST API routes with proper HTTP method mappings and middleware integration.
// Creates route definitions that connect HTTP endpoints to handlers with authentication and validation.
// Automatically registers routes in the main router for immediate API availability.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateRoute(name string) error {
	// Get project name from go.mod
	projectName, err := utils.GetModuleName()
	if err != nil {
		projectName = "your-project" // fallback
	}

	// Generate route file
	filename := fmt.Sprintf("route/%s.go", strings.ToLower(name))
	if err := utils.WriteFile(filename, templates.ResourceRouteTemplate, map[string]interface{}{
		"ProjectName": projectName,
		"Name":        name,
		"LowerName":   strings.ToLower(name),
	}); err != nil {
		return err
	}

	// Update v1.go to register the new routes
	return updateV1Routes(name)
}

// updateV1Routes automatically registers new resource routes in the main v1 router.
// Ensures new routes are immediately available without manual configuration.
func updateV1Routes(resourceName string) error {
	v1FilePath := "route/v1.go"

	// Check if v1.go exists
	if _, err := os.Stat(v1FilePath); os.IsNotExist(err) {
		return fmt.Errorf("v1.go file not found at %s. Make sure you're in the root directory of a Go To Oakhouse project", v1FilePath)
	}

	contentBytes, err := os.ReadFile(v1FilePath)
	if err != nil {
		return fmt.Errorf("failed to read v1.go: %v", err)
	}

	lines := strings.Split(string(contentBytes), "\n")
	setupFuncStart := -1
	setupFuncEnd := -1
	apiGroupLine := `api := app.Group("/api/v1")`
	apiGroupComment := "// API v1 routes"
	resourceComment := "// Setup resource routes"
	setupCall := fmt.Sprintf("Setup%sRoutes(api, db)", resourceName)

	// Find SetupRoutes function boundaries
	for i, line := range lines {
		if strings.Contains(line, "func SetupRoutes(") {
			setupFuncStart = i
		}
		if setupFuncStart != -1 && strings.TrimSpace(line) == "}" {
			setupFuncEnd = i
			break
		}
	}
	if setupFuncStart == -1 || setupFuncEnd == -1 {
		return fmt.Errorf("could not find SetupRoutes function in v1.go")
	}

	// Prevent duplicate SetupXxxRoutes call
	for _, line := range lines[setupFuncStart:setupFuncEnd] {
		if strings.TrimSpace(line) == setupCall {
			return nil // Already registered
		}
	}

	// Add `api := app.Group(...)` if missing
	apiGroupExists := false
	apiGroupIndex := -1
	for i := setupFuncStart; i < setupFuncEnd; i++ {
		if strings.TrimSpace(lines[i]) == apiGroupLine {
			apiGroupExists = true
			apiGroupIndex = i
			break
		}
	}

	if !apiGroupExists {
		// Insert API group before closing brace
		insertIndex := setupFuncEnd
		lines = append(lines[:insertIndex], append([]string{
			"\t" + apiGroupComment,
			"\t" + apiGroupLine,
		}, lines[insertIndex:]...)...)
		setupFuncEnd += 2
		apiGroupIndex = insertIndex + 1 // line number of api := ...
	}

	// Look for resource comment or SetupXxxRoutes block
	resourceCommentIndex := -1
	lastSetupCallIndex := -1
	for i := setupFuncStart; i < setupFuncEnd; i++ {
		line := strings.TrimSpace(lines[i])
		if line == resourceComment {
			resourceCommentIndex = i
		}
		if strings.HasPrefix(line, "Setup") && strings.Contains(line, "Routes(api, db)") {
			lastSetupCallIndex = i
		}
	}

	if resourceCommentIndex != -1 {
		// Insert below the comment
		insertIndex := resourceCommentIndex + 1
		lines = append(lines[:insertIndex], append([]string{"\t" + setupCall}, lines[insertIndex:]...)...)
	} else if lastSetupCallIndex != -1 {
		// Insert after last SetupXyzRoutes
		insertIndex := lastSetupCallIndex + 1
		lines = append(lines[:insertIndex], append([]string{"\t" + setupCall}, lines[insertIndex:]...)...)
	} else if apiGroupIndex != -1 {
		// Insert after API group
		insertIndex := apiGroupIndex + 1
		lines = append(lines[:insertIndex], append([]string{
			"",
			"\t" + resourceComment,
			"\t" + setupCall,
		}, lines[insertIndex:]...)...)
	} else {
		// Last fallback — just before end of function
		insertIndex := setupFuncEnd
		lines = append(lines[:insertIndex], append([]string{
			"",
			"\t" + apiGroupComment,
			"\t" + apiGroupLine,
			"",
			"\t" + resourceComment,
			"\t" + setupCall,
		}, lines[insertIndex:]...)...)
	}

	// Write final content
	return os.WriteFile(v1FilePath, []byte(strings.Join(lines, "\n")), 0644)
}
