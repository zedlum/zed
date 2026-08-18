package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/vmctl"
)

var hostDir string

// syncCmd configures (idempotently) a virtiofs share of hostDir into the
// dev VM. Requires virtiofsd already installed; never installs it.
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Configure a live virtiofs share into the dev VM",
	RunE: func(cmd *cobra.Command, args []string) error {
		return vmctl.SyncInit(domain, hostDir)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	wd, _ := os.Getwd()
	syncCmd.Flags().StringVarP(&hostDir, "host-dir", "p", wd, "host directory to share into the guest")
}
