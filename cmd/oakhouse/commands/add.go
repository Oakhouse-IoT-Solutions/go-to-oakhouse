// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"fmt"
	"os"

	"github.com/Oakhouse-Technology/go-to-oakhouse/cmd/oakhouse/generators"
	"github.com/spf13/cobra"
)

// AddCmd creates the 'add' command for adding features to existing projects.
// Provides functionality to add database support, authentication, middleware,
// and other features to enhance existing Oakhouse projects.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func AddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add features to your project",
		Long:  `Add various features like database support, authentication, middleware, and more to your existing Oakhouse project.`,
	}

	// Add subcommands
	cmd.AddCommand(addDatabaseCmd())

	return cmd
}

// addDatabaseCmd creates the command for adding database support to projects.
// Configures database connections, migrations, and GORM setup for data persistence
// with support for multiple database providers like PostgreSQL, MySQL, and SQLite.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func addDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Add database support to your project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := generators.AddDatabase(); err != nil {
				fmt.Fprintf(os.Stderr, "Error adding database: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("🚀 Database support added successfully!")
			fmt.Println("🏡 Proudly Created by Htet Waiyan From Oakhouse")
		},
	}
	return cmd
}