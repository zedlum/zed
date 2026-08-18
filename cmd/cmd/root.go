package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is zed's base command.
var rootCmd = &cobra.Command{
	Use:   "zed",
	Short: "Bootstrapping and dev tool for zedlum-os",
	Long:  "Clones component repos from the embedded manifest and drives the QEMU dev loop.",
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
