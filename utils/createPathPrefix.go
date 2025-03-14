package utils

import (
	"errors"
	"fmt"
	"os"
)

func CreatePrefix() (string, error) {

	// This is tested and works on WSL. It may or may not work on other systems.
	// This is tested and works on Windows. It may or may not work on other systems.
	// TODO: Add support for Linux
	// TODO: Add support for Mac

	if IsWSL() {
		home, err := os.UserHomeDir()

		if err != nil {
			return "", errors.New("error getting user home directory")
		}

		fmt.Println("Home directory: ", home)

		return home + "/", nil
	}

	if IsWindows() {
		home, err := os.Getwd()

		if err != nil {
			return "", errors.New("error getting user home directory")
		}

		fmt.Println("Home directory: ", home)

		return home + "/", nil
	}

	return "", errors.New("unsupported operating system")
}
