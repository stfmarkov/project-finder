package cmd

import (
	"pfinder/config"

	"github.com/spf13/cobra"
)

func searchForProjects(searchTerm string) ([]string, error) {
	projects, err := config.FindProject(searchTerm)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := searchForProjects(args[0])

		if err != nil {
			cmd.PrintErr(err)
			return
		}

		ChoiceSelector(projects, executeAction)
	},
}
