package config

import (
	"fmt"
	"os"
	"slices"
	"strings"

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

func AddProjectToConfig(project string) error {
	config, err := readFile()

	if err != nil {
		return err
	}

	if len(config.Projects) == 0 {
		config.Projects = []Project{}
	}

	isExistingProject := false

	for _, projectConfig := range config.Projects {
		if projectConfig.Path == project {
			isExistingProject = true
		}
	}

	if isExistingProject {
		return nil
	}

	config.Projects = append(config.Projects, Project{
		Path:     project,
		Alias:    strings.Split(project, "/")[len(strings.Split(project, "/"))-1],
		Commands: []string{},
	})

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
			Alias:    strings.Split(project, "/")[len(strings.Split(project, "/"))-1],
			Commands: []string{command},
		}

		config.Projects = append(config.Projects, projectConfig)
	}

	return UpdateFile(config)
}

func DeleteCommandForProject(project string, command string) {
	config, err := readFile()

	if err != nil {
		return
	}

	for i, projectConfig := range config.Projects {
		if projectConfig.Path == project {
			for j, projectCommand := range projectConfig.Commands {
				if projectCommand == command {
					config.Projects[i].Commands = append(projectConfig.Commands[:j], projectConfig.Commands[j+1:]...)
				}
			}
		}
	}

	UpdateFile(config)
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

func GetProjects() ([]string, error) {
	config, err := readFile()

	if err != nil {
		return nil, err
	}

	paths := []string{}

	for _, project := range config.Projects {
		paths = append(paths, project.Path)
	}

	return paths, nil
}

func FindProject(searchTerm string) ([]string, error) {
	config, err := readFile()

	if err != nil {
		return nil, err
	}

	projects := []string{}

	for _, project := range config.Projects {
		if strings.Contains(project.Path, searchTerm) {
			projects = append(projects, project.Path)
		}
	}

	return projects, nil
}

func SaveProjects(projects []string) error {
	for _, project := range projects {
		err := AddProjectToConfig(project)

		if err != nil {
			return err
		}
	}

	return nil
}
