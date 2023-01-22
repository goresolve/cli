package utils

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func LogMessage(msg string) {
	yellow := color.New(color.FgHiBlue).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	fmt.Printf("%s %s", yellow("[Resolve CLI]"), white(msg))
}

func Message(tag, msg string) {
	yellow := color.New(color.FgHiBlue).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	fmt.Printf("%s %s", yellow(tag), white(msg))
}

func ErrorMessage(msg string) {
	red := color.New(color.FgHiRed).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	fmt.Printf("%s %s", red("[Resolve CLI]"), white(msg))
	os.Exit(0)
}
