package cmd

import (
	"os"
	"strings"
	"text/template"
)

func CreateController(name string) {
	if !isDirectoryExists("internal") {
		ErrorMessage("internal folder not found. Please run 'resolve create' command first.")
	}

	if !isDirectoryExists("internal/controllers") {
		ErrorMessage("internal/controllers folder not found. Please run 'resolve create' command first.")
	}

	if isDirectoryExists("internal/controllers/" + name) {
		ErrorMessage("Controller already exists.")
	}

	err := os.Mkdir("internal/controllers/"+pathName(name), 0777)
	if err != nil {
		ErrorMessage("Error while creating folder: " + err.Error())
	} else {
		LogMessage("Folder " + pathName(name) + " created successfully.\n")
	}

	createService(name)
	createController(name)
}

func isDirectoryExists(name string) bool {
	exist, err := os.Stat(name)
	if err != nil {
		return false
	}

	return exist.IsDir()
}

func pathName(name string) string {
	return strings.ToLower(name) + "_controller"
}

func fileName(name string) string {
	return strings.ToLower(name)
}

func createService(name string) {
	tmpl, err := template.New("service").Parse(ServiceScheme)
	if err != nil {
		ErrorMessage("Error while creating template: " + err.Error())
	}

	file, err := os.Create("internal/controllers/" + pathName(name) + "/" + fileName(name) + ".service.go")
	if err != nil {
		ErrorMessage("Error while creating file: " + err.Error())
	} else {
		LogMessage("File " + fileName(name) + ".service.go created successfully.\n")
	}

	err = tmpl.Execute(file, map[string]string{"Name": fileName(name)})
}

func createController(name string) {
	tmpl, err := template.New("controller").Parse(ControllerScheme)
	if err != nil {
		ErrorMessage("Error while creating template: " + err.Error())
	}

	file, err := os.Create("internal/controllers/" + pathName(name) + "/" + fileName(name) + ".controller.go")
	if err != nil {
		ErrorMessage("Error while creating file: " + err.Error())
	} else {
		LogMessage("File " + fileName(name) + ".controller.go created successfully.\n")
	}

	err = tmpl.Execute(file, map[string]string{"Name": fileName(name)})
}
