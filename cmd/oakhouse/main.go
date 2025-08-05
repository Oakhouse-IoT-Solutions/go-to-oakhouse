// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"fmt"
	"os"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/commands"
	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/utils"
	"github.com/spf13/cobra"
)

// main initializes and executes the Oakhouse CLI application with all available commands.
// Sets up the root command with version information and registers all subcommands
// for project creation, code generation, feature addition, serving, and building.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func main() {
	rootCmd := &cobra.Command{
		Use:   "oakhouse",
		Short: "🚀 Go To Oakhouse - A powerful Go framework for rapid API development",
		Long: `🚀 Go To Oakhouse is a Go framework for building APIs with clean architecture patterns.

🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡

Features:
• Clean Architecture patterns
• Rapid API development
• Auto-generated CRUD operations
• Built-in authentication middleware
• Database support
• Docker containerization`,
		Version: utils.Version,
	}

	// Add subcommands from commands package
	rootCmd.AddCommand(commands.NewCmd())
	rootCmd.AddCommand(commands.GenerateCmd())
	rootCmd.AddCommand(commands.IntegrateCmd())
	rootCmd.AddCommand(commands.ServeCmd())
	rootCmd.AddCommand(commands.BuildCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

//
