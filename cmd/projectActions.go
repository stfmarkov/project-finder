package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func showProjectActions(project string) {
	fmt.Println("Will show project actions for", project)

	actions := []string{
		"Navigate to project",
		// TODO: Add more actions
	}

	implementedActions := map[string]func(string){
		"Navigate to project": func(project string) {
			fmt.Println("Navigating to project", project)

			project = strings.ReplaceAll(project, "/home/", "")

			cmd := exec.Command("cmd.exe", "/c", "start", "wsl.exe", "bash", "-c", "cd ~/../"+project+" && exec bash")

			err := cmd.Run()

			if err != nil {
				fmt.Println("Error navigating to project", err)
			}
		},
	}

	takeAction := func(action string) {
		implementedActions[action](project)
	}

	ChoiceSelector(actions, takeAction)
}
