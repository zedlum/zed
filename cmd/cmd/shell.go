package cmd

import (
	"os/user"

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
	def := ""
	if u, err := user.Current(); err == nil {
		def = u.Username
	}
	shellCmd.Flags().StringVarP(&shellUser, "user", "u", def, "SSH user on the guest")
}
