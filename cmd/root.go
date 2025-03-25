package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pfinder",
	Short: "A tool to switch projects and manage notes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to pFinder! Use --help for commands.")
	},
}

// Execute runs the CLI
func Execute() {

	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(findCmd)
	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
