package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func createMainFile() {
	var lines = []string{
		"package main",
		"",
		"import (",
		"    \"fmt\"",
		"    \"log\"",
		")",
		"",
		"func main() {",
		"    var i string",
		"    _, err := fmt.Scan(&i)",
		"    if err != nil {",
		"        log.Fatal(err)",
		"    }",
		"}",
	}

	if f, err := os.Create("main.go"); err != nil {
		fmt.Println("Cannot create main file: ", err)
		os.Exit(1)
	} else {
		defer f.Close()
		for _, line := range lines {
			_, err := f.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func getProjectName() string {
	cmdInput := os.Args[1:]
	var name string
	if len(cmdInput) == 0 {
		fmt.Println("No name provided. Using current directory as name!")

		wd, err := os.Getwd()

		if err != nil {
			fmt.Println("Cannot get working directory. ", err)
			os.Exit(1)
		}
		dirs := strings.Split(wd, "/")
		name = dirs[len(dirs)-1]
	} else {
		name = cmdInput[0]
	}
	return name
}

func initGoProject(projectName string) {
	cmd := exec.Command("go", "mod", "init", projectName)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error initializing project: ", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Initializing Go project...")
	projectName := getProjectName()
	mod_name := fmt.Sprintf("kunniii/codeforces/%s", projectName)
	initGoProject(mod_name)
	fmt.Println("Project initialized successfully!")
	createMainFile()
	fmt.Println("Create main.go!")
}
