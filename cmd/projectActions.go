package cmd

import "fmt"

func showProjectActions(project string) {
	fmt.Println("Will show project actions for", project)

	actions := []string{
		"Navigate to project",
		// TODO: Add more actions
	}

	implementedActions := map[string]func(string){
		"Navigate to project": func(project string) {
			fmt.Println("Navigating to project", project)
		},
	}

	takeAction := func(action string) {
		implementedActions[action](project)
	}

	ChoiceSelector(actions, takeAction)
}
