package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zedlum/zed/internal/build"
	"github.com/zedlum/zed/internal/manifest"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build every manifest repo that has a Makefile",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := manifest.Load()
		if err != nil {
			return err
		}
		return build.Components(m)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
