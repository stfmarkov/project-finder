package utils

import (
	"errors"
	"os"
)

func CreatePrefix() (string, error) {
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
