package config

import (
	"fmt"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

const configFileName = "pFinder.yaml"

type Config struct {
	ProjectDirs []string
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

func AddProjectDir(projectDir string) error {
	config, err := readFile()

	if err != nil {
		return err
	}

	if slices.Contains(config.ProjectDirs, projectDir) {
		return fmt.Errorf("project directory already exists")
	}

	config.ProjectDirs = append(config.ProjectDirs, projectDir)

	yamlConfig, err := yaml.Marshal(config)

	if err != nil {
		return fmt.Errorf("error adding project directory: %w", err)
	}

	err = os.WriteFile(configFileName, yamlConfig, 0644)

	if err != nil {
		return fmt.Errorf("error adding project directory: %w", err)
	}

	return nil
}
