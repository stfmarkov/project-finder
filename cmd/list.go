package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var projects []string = []string{}

var projectDirs []string = []string{
	"dev", "projects", "workspace", "work", "code", "repos", "repositories",
}

func findProjects(path string, projects []string) []string {

	fmt.Println(path)

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(errors.New("error reading directory"))
	}

	for _, file := range files {

		fmt.Println(file)

		if file.IsDir() {

			fmt.Println(file)

			if file.Name() == ".git" {
				projects = append(projects, path)
				return projects
			}

			// findProjects(path+"/"+file.Name(), projects)
		}
	}

	return projects
}

func createPrefix() string {
	home, err := os.UserHomeDir()

	if err != nil {
		return home
	}

	return "/"
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Listing all projects...")

		projects = findProjects(createPrefix()+"home", projects)

		for _, projectDir := range projectDirs {
			fmt.Println(projectDir, "projects:")
		}
	},
}
