// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package commands

import (
	"fmt"
	"os"

	"github.com/Oakhouse-IoT-Solutions/go-to-oakhouse/cmd/oakhouse/generators"
	"github.com/spf13/cobra"
)

// ServeCmd creates the command for starting the development server.
// Provides hot reload capabilities using Air if available, with configurable host and port,
// enabling live development experience with automatic server restart on code changes.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func ServeCmd() *cobra.Command {
	var port int
	var host string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the development server with hot reload",
		Run: func(cmd *cobra.Command, args []string) {
			if err := generators.StartDevServer(host, port); err != nil {
				fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	cmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to bind the server to")

	return cmd
}
