package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/build"
)

var baseISO string

// isoCmd launches Cubic for the manual ISO repack step. Not scripted yet:
// see internal/build.LaunchCubic.
var isoCmd = &cobra.Command{
	Use:   "iso",
	Short: "Launch Cubic to repack the base ISO (manual step, v1)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return build.LaunchCubic(baseISO)
	},
}

func init() {
	rootCmd.AddCommand(isoCmd)
	isoCmd.Flags().StringVar(&baseISO, "base-iso", "", "path to the base Ubuntu Budgie ISO")
}
