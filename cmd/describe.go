package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe subject in the subcommand",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please specify the subject to be described")
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
