package generator

import (
	"fmt"
	"github.com/goresolve/cli/resolve/schemes"
	"github.com/goresolve/cli/resolve/utils"
	"os"
	"os/exec"
)

func createDir() string {
	utils.LogMessage("Enter folder name to create project: ")

	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		utils.ErrorMessage("Error while reading folder name: " + err.Error())
	}

	err = os.Mkdir(path, 0777)
	if err != nil {
		utils.ErrorMessage("Error while creating folder: " + err.Error())
	} else {
		utils.LogMessage("Folder created successfully.\n")
	}

	finalPath, err := os.Getwd()
	if err != nil {
		utils.ErrorMessage("Error while getting current directory: " + err.Error())
	}

	return finalPath + "\\" + path
}

func createMod(path string) {
	utils.LogMessage("Enter module name to create go.mod file: ")

	var module string
	_, err := fmt.Scanln(&module)
	if err != nil {
		utils.ErrorMessage("Error while reading folder name: " + err.Error())
	}

	modCreate := exec.Command("go", "mod", "init", module)
	modCreate.Dir = path
	err = modCreate.Run()
	if err != nil {
		utils.ErrorMessage("Error while creating go.mod file: " + err.Error())
	} else {
		utils.LogMessage("go.mod file created successfully.\n")
	}
}

func installDependencies(path string) {
	getPackage := exec.Command("go", "get", "github.com/goresolve/resolve")
	getPackage.Dir = path
	err := getPackage.Run()
	if err != nil {
		utils.ErrorMessage("Error while installing dependencies: " + err.Error())
	} else {
		utils.LogMessage("Dependencies installed successfully.\n")
	}
}

func createDirectories(path string) {
	folders := []string{"cmd", "internal", "internal\\controllers", "internal\\middlewares", "web", "web\\templates", "web\\static"}
	for _, folder := range folders {
		err := os.Mkdir(path+"\\"+folder, 0777)
		if err != nil {
			utils.ErrorMessage("Error while creating folder: " + err.Error())
		} else {
			utils.LogMessage("Folder " + folder + " created successfully.\n")
		}
	}
}

func createFiles(path string) {
	file, _ := os.Create(path + "\\cmd\\main.go")
	_, err := file.Write([]byte(schemes.BaseApplication))
	if err != nil {
		utils.ErrorMessage("Error while creating file: " + err.Error())
	}
}

func CreateProject() {
	path := createDir()
	createMod(path)
	installDependencies(path)
	createDirectories(path)
	createFiles(path)
}
