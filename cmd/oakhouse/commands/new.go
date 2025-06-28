// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"fmt"
	"os"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/generators"
	"github.com/spf13/cobra"
)

// NewCmd creates the 'new' command for generating new Oakhouse projects.
// Initializes a complete project structure with all necessary files, dependencies,
// and configuration for rapid API development with clean architecture patterns.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new [project-name]",
		Short: "Create a new Go To Oakhouse project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			if err := generators.CreateNewProject(projectName); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating project: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("\n🚀 Project '%s' created successfully!\n", projectName)
			fmt.Printf("🏡 Proudly Created by Htet Waiyan From Oakhouse\n\n")
			fmt.Printf("Next steps:\n")
			fmt.Printf("  cd %s\n", projectName)
			fmt.Printf("  cp .env.example .env\n")
			fmt.Printf("  oakhouse serve\n\n")
			fmt.Printf("🚀 Happy coding with Oakhouse! 🏡\n")
		},
	}
	return cmd
}