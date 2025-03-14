package main

import (
	"pfinder/cmd"
	"pfinder/config"
)

func main() {
	config.CreateMissConfig()
	cmd.Execute()
}
