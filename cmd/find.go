package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Finding project...")
	},
}
