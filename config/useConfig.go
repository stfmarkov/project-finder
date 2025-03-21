package config

import (
	"fmt"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

const configFileName = "pFinder.yaml"

type Project struct {
	Path     string
	Alias    string
	Commands []string
}

type Config struct {
	ProjectDirs []string
	Projects    []Project
	// ProjectAliases  map[string]string
	// ProjectCommands map[string][]string
}

func CreateMissConfig() {

	_, err := os.ReadFile(configFileName)

	if err == nil {
		return
	}

	config := Config{
		ProjectDirs: []string{
			"dev", "projects", "workspace", "work", "code", "repos", "repositories",
		},
	}

	yamlConfig, err := yaml.Marshal(config)

	if err != nil {
		panic(err)
	}

	os.WriteFile(configFileName, yamlConfig, 0644)
}

func readFile() (Config, error) {
	var config Config

	configData, err := os.ReadFile(configFileName)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	return config, nil
}

func GetProjectDirs() ([]string, error) {
	config, err := readFile()

	if err != nil {
		return nil, err
	}

	return config.ProjectDirs, nil
}

func UpdateFile(config Config) error {
	yamlConfig, err := yaml.Marshal(config)

	if err != nil {
		return fmt.Errorf("error updating config file: %w", err)
	}

	err = os.WriteFile(configFileName, yamlConfig, 0644)

	if err != nil {
		return fmt.Errorf("error updating config file: %w", err)
	}

	return nil
}

func AddProjectDir(projectDir string) error {
	config, err := readFile()

	if err != nil {
		return err
	}

	if slices.Contains(config.ProjectDirs, projectDir) {
		return fmt.Errorf("project directory already exists")
	}

	config.ProjectDirs = append(config.ProjectDirs, projectDir)

	return UpdateFile(config)
}

func AddCommandForProject(project string, command string) error {
	config, err := readFile()

	if err != nil {
		return err
	}

	if len(config.Projects) == 0 {
		config.Projects = []Project{}
	}

	isExistingProject := false

	for i, projectConfig := range config.Projects {
		if projectConfig.Path == project {
			config.Projects[i].Commands = append(config.Projects[i].Commands, command)
			isExistingProject = true
		}
	}

	if !isExistingProject {
		projectConfig := Project{
			Path:     project,
			Alias:    "",
			Commands: []string{command},
		}

		config.Projects = append(config.Projects, projectConfig)
	}

	return UpdateFile(config)
}

func GetCommandsForProject(project string) ([]string, error) {
	config, err := readFile()

	if err != nil {
		return nil, err
	}

	for _, projectConfig := range config.Projects {
		if projectConfig.Path == project {
			return projectConfig.Commands, nil
		}
	}

	return nil, fmt.Errorf("project not found")
}
