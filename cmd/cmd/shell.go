package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/vmctl"
)

var shellUser string

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open an SSH session into the dev VM",
	RunE: func(cmd *cobra.Command, args []string) error {
		return vmctl.Shell(domain, shellUser)
	},
}

func init() {
	vmCmd.AddCommand(shellCmd)
	// guest account is always "zed", never the host's own username
	shellCmd.Flags().StringVarP(&shellUser, "user", "u", "zed", "SSH user on the guest")
}
