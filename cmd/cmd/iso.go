package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// isoCmd represents the iso command
var isoCmd = &cobra.Command{
	Use:   "iso",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("iso called")
	},
}

func init() {
	rootCmd.AddCommand(isoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
