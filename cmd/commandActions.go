package cmd

import (
	"fmt"
	"pfinder/config"
	"strings"
)

const projectCommandsSplitter = "{project.command}"

func deleteCommand(projectCommand string) {
	project, command := strings.Split(projectCommand, projectCommandsSplitter)[0], strings.Split(projectCommand, projectCommandsSplitter)[1]

	config.DeleteCommandForProject(project, command)
}

func showCommandActions(project string, command string) {
	fmt.Println("Will show command actions for", command)

	actions := []string{
		"Delete command",
	}

	implementedActions := map[string]func(string){
		"Delete command": deleteCommand,
	}

	takeAction := func(action string) {
		implementedActions[action](fmt.Sprintf("%v%s%v", project, projectCommandsSplitter, command))
	}

	ChoiceSelector(actions, takeAction)
}
