package cmd

import (
	"fmt"
	"os/exec"
	"pfinder/utils"
	"strings"
)

const AddDirActionStr = "These are not the projects I'm looking for"

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

	cmd := exec.Command("cmd.exe", "/c", "start", "wsl.exe", "bash", "-c", "code ~/../"+project)

	err := cmd.Start()

	if err != nil {
		fmt.Println("Error opening project", err)
	}
}

func openInWindows(project string) {
	cmd := exec.Command("cmd.exe", "/c", "start", "cmd", "/k", "code", project)

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

func showProjectActions(project string) {
	fmt.Println("Will show project actions for", project)

	actions := []string{
		"Navigate to project",
		"Open project",
		// TODO: Add more actions
	}

	implementedActions := map[string]func(string){
		"Navigate to project": navigateToProject,
		"Open project":        openProject,
	}

	takeAction := func(action string) {
		implementedActions[action](project)
	}

	ChoiceSelector(actions, takeAction)
}
