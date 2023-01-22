package main

import (
	"github.com/goresolve/cli/cmd"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		cmd.Message("[Resolve CLI]", "Command not found. List of available commands:\n")
		cmd.Message("[Resolve CLI]", "resolve create | Project creation wizard. Sets the dependencies and creates folders.\n")
		cmd.Message("[Resolve CLI]", "resolve generate controller <name> | Creates a controller with the specified name.\n")
		return
	}

	if args[1] == "create" {
		cmd.CreateProject()
	}

	if (args[1] == "generate" || args[1] == "g") && (args[2] == "controller" || args[2] == "co") {
		cmd.CreateController(args[3])
	}
}
