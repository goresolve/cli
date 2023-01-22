package main

import (
	"github.com/goresolve/resolve"
)

func main() {
	app := resolve.Setup()

	app.RegisterTemplates(
		"./web/templates", ".html", true,
	)

	app.Run(":8080")
}
