package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/component"
	"github.com/zedlum/zed/internal/manifest"
)

// componentCmd manages the manifest's repos.
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Manage component repos from the embedded manifest",
}

var componentSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Clone missing repos and check out each to its pinned version",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := manifest.Load()
		if err != nil {
			return err
		}
		return component.Sync(m)
	},
}

var componentStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report each repo's sync state against the manifest",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := manifest.Load()
		if err != nil {
			return err
		}
		statuses, err := component.Check(m)
		if err != nil {
			return err
		}
		if len(statuses) == 0 {
			fmt.Println("no repos in manifest")
			return nil
		}
		diverged := false
		for _, s := range statuses {
			fmt.Printf("%-30s %-10s %s\n", s.Repo.Path, s.State, s.HeadRef)
			if s.State != "clean" {
				diverged = true
			}
		}
		if diverged {
			os.Exit(1)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(componentCmd)
	componentCmd.AddCommand(componentSyncCmd)
	componentCmd.AddCommand(componentStatusCmd)
}
