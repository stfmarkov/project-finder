package cmd

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var projects []string = []string{}

var projectDirs []string = []string{
	"dev", "projects", "workspace", "work", "code", "repos", "repositories", "boot-dev",
}

func findProjects(path string, projects *[]string) {
	files, err := os.ReadDir(path)
	if err != nil {
		// fmt.Println(errors.New("error reading directory"))
	}

	for _, file := range files {
		if file.IsDir() {
			if file.Name() == ".git" && !slices.Contains(*projects, path) {
				*projects = append(*projects, path)
				return
			}

			findProjects(path+"/"+file.Name(), projects)
		}
	}
}

func createPrefix() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", errors.New("error getting user home directory")
	}

	// This is tested and works on WSL. It may or may not work on other systems.
	// TODO: Add support for Windows
	// TODO: Add support for Linux
	// TODO: Add support for Mac

	return home + "/", nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		prefix, err := createPrefix()

		if err != nil {
			fmt.Println(errors.New("error getting user home directory"))
			return
		}

		for _, projectDir := range projectDirs {
			// TODO: This can be done concurrently
			findProjects(prefix+projectDir, &projects)
		}

		ChoiceSelector(projects)
	},
}
