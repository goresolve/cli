package main

import (
	"github.com/goresolve/cli/resolve/generator"
	"github.com/goresolve/cli/resolve/utils"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		utils.Message("[Resolve CLI]", "Command not found. List of available commands:\n")
		utils.Message("[Resolve CLI]", "resolve create | Project creation wizard. Sets the dependencies and creates folders.\n")
		return
	}

	if args[1] == "create" {
		generator.CreateProject()
	}
}
