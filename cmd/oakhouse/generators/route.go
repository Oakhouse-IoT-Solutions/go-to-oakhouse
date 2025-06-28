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
		return fmt.Errorf("v1.go file not found at %s. Make sure you are in the root directory of a Go To Oakhouse project", v1FilePath)
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
			// If neither comment exists, add the comment before the closing brace
			closeBraceIndex := strings.LastIndex(contentStr, "}")
			if closeBraceIndex == -1 {
				return fmt.Errorf("could not find insertion point in v1.go: malformed file structure")
			}

			// Insert the comment and route call before the closing brace
			newContent := contentStr[:closeBraceIndex] + "\n\t" + commentPattern + "\n\t" + setupCall + "\n" + contentStr[closeBraceIndex:]
			return os.WriteFile(v1FilePath, []byte(newContent), 0644)
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