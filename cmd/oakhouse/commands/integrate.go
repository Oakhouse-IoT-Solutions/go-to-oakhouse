// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"fmt"
	"os"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/generators"
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
		Long:  `Integrate various features like database support, authentication, middleware, and more to your existing Oakhouse project.`,
	}

	// Add subcommands
	cmd.AddCommand(integrateDatabaseCmd())

	return cmd
}

// integrateDatabaseCmd creates the command for integrating database support to projects.
// Configures database connections, migrations, and GORM setup for data persistence
// with support for multiple database providers like PostgreSQL, MySQL, and SQLite.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func integrateDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Integrate database support to your project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := generators.IntegrateDatabase(); err != nil {
				fmt.Fprintf(os.Stderr, "Error integrating database: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("🚀 Database support integrated successfully!")
			fmt.Println("🏡 Proudly Created by Htet Waiyan From Oakhouse")
		},
	}
	return cmd
}
