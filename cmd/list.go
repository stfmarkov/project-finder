package cmd

import (
	"fmt"
	"pfinder/config"

	"github.com/spf13/cobra"
)

func listAllProjects() error {
	var err error
	projects, err := config.GetProjects()

	if err == nil {
		ChoiceSelector(append(projects, AddDirActionStr), executeAction)
	}

	return err
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all fetched projects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("If the project you are looking for is not listed you may need to fetch the projects first. ( with 'pFinder fetch' ) ")
		err := listAllProjects()

		if err != nil {
			fmt.Println("Error listing projects", err)
		}
	},
}
