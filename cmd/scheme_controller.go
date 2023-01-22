package cmd

const ServiceScheme = `package {{.Name}}

import "github.com/goresolve/resolve"

func get(c *resolve.Ctx) {
	// implement your logic here
}

func create(c *resolve.Ctx) {
	// implement your logic here
}

func update(c *resolve.Ctx) {
	// implement your logic here
}

func delete(c *resolve.Ctx) {
	// implement your logic here
}`

const ControllerScheme = `package {{.ControllerName}}

import "github.com/goresolve/resolve"

func InitializeController() resolve.Controller {
	controller := resolve.NewController("/{{.Name}}")

	controller.Get("/", get)
	controller.Post("/", create)
	controller.Put("/", update)
	controller.Delete("/", delete)

	return controller
}
`
