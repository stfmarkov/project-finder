package cmd

import (
	"fmt"
	"pfinder/config"

	"github.com/spf13/cobra"
)

func executeAction(project string) {
	if project == AddDirActionStr {
		DirectInput("Enter the path to the project: ", func(project string) {

			err := config.AddProjectDir(project)

			if err != nil {
				fmt.Println("Error adding project directory:", err)
				return
			}

			fmt.Println("Project directory added successfully")
			listAllProjects()

		})
	} else {
		showProjectActions(project)
	}
}

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
	Short: "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		err := listAllProjects()

		if err != nil {
			fmt.Println("Error listing projects", err)
		}

		fmt.Println("If the project you are looking for is not listed you may need to fetch the projects first. ( with 'pFinder fetch' ) ")
	},
}
