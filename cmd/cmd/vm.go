package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

const defaultDomain = "zedlum-dev"

var domain string

// vmCmd groups operations on the dev VM. zed never creates the domain
// itself — it must already exist (virt-install / virt-manager, done by hand).
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Manage the dev VM (must already exist)",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := exec.Command("virsh", "domstate", domain).CombinedOutput()
		fmt.Print(string(out))
		return err
	},
}

func init() {
	rootCmd.AddCommand(vmCmd)
	vmCmd.PersistentFlags().StringVarP(&domain, "domain", "d", defaultDomain, "libvirt domain name")
}
