package cmd

import (
	"fmt"
	"os/exec"
	"pfinder/config"
	"pfinder/utils"
	"strings"
)

const AddDirActionStr = "These are not the projects I'm looking for"

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

func navigateInWsl(project string) {
	project = strings.ReplaceAll(project, "/home/", "")

	cmd := exec.Command("cmd.exe", "/c", "start", "wsl.exe", "bash", "-c", "cd ~/../"+project+" && exec bash")

	err := cmd.Start()

	if err != nil {
		fmt.Println("Error navigating to project", err)
	}
}

func navigateInWindows(project string) {
	cmd := exec.Command("cmd.exe", "/c", "start", "cmd", "/k", "cd", project)

	err := cmd.Start()

	if err != nil {
		fmt.Println("Error navigating to project", err)
	}
}

func navigateToProject(project string) {
	fmt.Println("Navigating to project", project)

	if utils.IsWSL() {
		navigateInWsl(project)
		return
	}

	if utils.IsWindows() {
		navigateInWindows(project)
		return
	}

	fmt.Println("Unsupported OS")
}

func openInWsl(project string) {
	project = strings.ReplaceAll(project, "/home/", "")

	cmd := exec.Command("cmd.exe", "/c", "start", "wsl.exe", "bash", "-c", "cd ~/../"+project+" && code . && exec bash")

	err := cmd.Start()

	if err != nil {
		fmt.Println("Error opening project", err)
	}
}

func openInWindows(project string) {
	cmd := exec.Command("cmd.exe", "/c", "start", "cmd", "/k", "cd /d"+project+" && code .")

	err := cmd.Start()

	if err != nil {
		fmt.Println("Error opening project", err)
	}
}

func openProject(project string) {
	fmt.Println("Navigating to project", project)

	if utils.IsWSL() {
		openInWsl(project)
		return
	}

	if utils.IsWindows() {
		openInWindows(project)
		return
	}

	fmt.Println("Unsupported OS")
}

func runInWsl(project string, commands string) {
	project = strings.ReplaceAll(project, "/home/", "")

	cmd := exec.Command("cmd.exe", "/c", "start", "wsl.exe", "bash", "-c", "code ~/../"+project)

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error running commands for project", err)
		return
	}

	fmt.Println("Running commands for project", "cd ~/../"+project+" && "+commands+" && exec bash")

	cmd = exec.Command("cmd.exe", "/c", "start", "wsl.exe", "bash", "-c", "cd ~/../"+project+" && "+commands+"; read -p 'Press enter to exit...'")

	err = cmd.Start()

	if err != nil {
		fmt.Println("Error running commands for project", err)
	}

}

func runInWindows(project string, commands string) {
	cmd := exec.Command("cmd.exe", "/c", "start", "cmd", "/k", "cd /d"+project+" && code . && "+commands)

	err := cmd.Start()

	if err != nil {
		fmt.Println("Error running commands for project", err)
	}

}

func runProjectCommands(project string) {

	commands, err := config.GetCommandsForProject(project)

	if err != nil {
		fmt.Println("Error getting commands for project", err)
		return
	}

	commandsString := strings.Join(commands, " && ")

	if utils.IsWSL() {
		runInWsl(project, commandsString)
		return
	}

	if utils.IsWindows() {
		runInWindows(project, commandsString)
		return
	}
}

func addCommand(project string) {
	DirectInput("Adding command for project ( If the commands you entered do not work ... skill issues ): ", func(command string) {

		fmt.Println("Will add command for project", command)

		if command == "" {
			fmt.Println("No command entered")
			return
		}

		config.AddCommandForProject(project, command)
	})
}

func showCustomCommands(project string) {
	commands, err := config.GetCommandsForProject(project)

	if err != nil {
		fmt.Println("Error getting commands for project", err)
		return
	}

	ChoiceSelector(commands, func(command string) {
		showCommandActions(project, command)
	})
}

func showProjectActions(project string) {
	fmt.Println("Will show project actions for", project)

	actions := []string{
		"Run project commands",
		"Navigate to project",
		"Open project",
		"Add command",
		"Show custom commands",
		// TODO: Add more actions
	}

	implementedActions := map[string]func(string){
		"Run project commands": runProjectCommands,
		"Navigate to project":  navigateToProject,
		"Open project":         openProject,
		"Add command":          addCommand,
		"Show custom commands": showCustomCommands,
	}

	takeAction := func(action string) {
		implementedActions[action](project)
	}

	ChoiceSelector(actions, takeAction)
}
