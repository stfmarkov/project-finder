package cmd

import (
	"errors"
	"fmt"
	"os"
	"pfinder/config"
	"pfinder/utils"
	"slices"

	"github.com/spf13/cobra"
)

var projects []string = []string{}

func findProjects(path string, projects *[]string) {
	files, err := os.ReadDir(path)
	if err != nil {
		// fmt.Println(errors.New("error reading directory"))
	}

	for _, file := range files {
		if file.IsDir() {
			if (file.Name() == ".git" || file.Name() == "node_modules") && !slices.Contains(*projects, path) {
				*projects = append(*projects, path)
				return
			}

			findProjects(path+"/"+file.Name(), projects)
		}
	}
}

func fetchProjects() {
	prefix, err := utils.CreatePrefix()

	if err != nil {
		fmt.Println(errors.New("error getting user home directory"))
		return
	}

	var projectDirs, fileErr = config.GetProjectDirs()

	if fileErr != nil {
		fmt.Println(errors.New("error reading config file"))
		return
	}

	for _, projectDir := range projectDirs {
		// TODO: This can be done concurrently
		findProjects(prefix+projectDir, &projects)
	}

	config.SaveProjects(projects)

	if err != nil {
		fmt.Println(errors.New("error listing projects"))
	}
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches and saves all projects",
	Run: func(cmd *cobra.Command, args []string) {
		fetchProjects()
		listAllProjects()
	},
}
