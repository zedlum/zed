package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/vmctl"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Gracefully shut down the dev VM",
	RunE: func(cmd *cobra.Command, args []string) error {
		return vmctl.Stop(domain)
	},
}

func init() {
	vmCmd.AddCommand(stopCmd)
}
