// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"fmt"
	"os"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/generators"
	"github.com/spf13/cobra"
)

// BuildCmd creates the command for building the application for production.
// Compiles the application into an optimized binary executable ready for deployment,
// with all dependencies statically linked for production environments.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func BuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the application for production",
		Run: func(cmd *cobra.Command, args []string) {
			if err := generators.BuildApplication(); err != nil {
				fmt.Fprintf(os.Stderr, "Error building application: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Application built successfully!")
		},
	}
	return cmd
}
