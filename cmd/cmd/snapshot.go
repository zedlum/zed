package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/vmctl"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Manage dev VM snapshots",
}

var snapshotCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a named snapshot",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return vmctl.SnapshotCreate(domain, args[0])
	},
}

var snapshotRevertCmd = &cobra.Command{
	Use:   "revert <name>",
	Short: "Revert to a named snapshot",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return vmctl.SnapshotRevert(domain, args[0])
	},
}

func init() {
	vmCmd.AddCommand(snapshotCmd)
	snapshotCmd.AddCommand(snapshotCreateCmd)
	snapshotCmd.AddCommand(snapshotRevertCmd)
}
