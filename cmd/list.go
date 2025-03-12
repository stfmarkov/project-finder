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
	fmt.Println(path)

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(errors.New("error reading directory"))
	}

	for _, file := range files {

		fmt.Println(file)

		if file.IsDir() {

			fmt.Println(file.Name(), "File name")

			if file.Name() == ".git" && !slices.Contains(*projects, path) {
				*projects = append(*projects, path)

				fmt.Println(projects)

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
		fmt.Println("Listing all projects...")

		prefix, err := createPrefix()

		if err != nil {
			fmt.Println(errors.New("error getting user home directory"))
			return
		}

		for _, projectDir := range projectDirs {
			findProjects(prefix+projectDir, &projects)
		}

		fmt.Println(projects, "Final state")

	},
}
