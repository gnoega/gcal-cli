/*
Copyright © 2024 Agung Firmansyah gnoega@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const version string = "v0.6.1"

var rootCmd = &cobra.Command{
	Use:     "gcal-cli",
	Short:   "Google calendar cli tool",
	Version: version,
	Long: `
 ██████╗  ██████╗ █████╗ ██╗       ██████╗██╗     ██╗
██╔════╝ ██╔════╝██╔══██╗██║      ██╔════╝██║     ██║
██║  ███╗██║     ███████║██║█████╗██║     ██║     ██║
██║   ██║██║     ██╔══██║██║╚════╝██║     ██║     ██║
╚██████╔╝╚██████╗██║  ██║███████╗ ╚██████╗███████╗██║
 ╚═════╝  ╚═════╝╚═╝  ╚═╝╚══════╝  ╚═════╝╚══════╝╚═╝
                                                     
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
