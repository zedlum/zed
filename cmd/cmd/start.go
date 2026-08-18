package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/vmctl"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the dev VM",
	RunE: func(cmd *cobra.Command, args []string) error {
		return vmctl.Start(domain)
	},
}

func init() {
	vmCmd.AddCommand(startCmd)
}
