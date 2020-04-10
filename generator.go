package main

import (
	"fmt"
	"github.com/art-injener/QtProjectGenerator/template"
	"github.com/art-injener/QtProjectGenerator/tree"
	"os"
)

func main() {
	fmt.Println("PROJECT GENERATION START")

	if !(len(os.Args) == 2) {
		fmt.Println("usage : ./generator project_name. \nPROJECT GENERATION ABORT")
		os.Exit(0)
	}
	projectName := os.Args[1]
	var path string
	if currentDir, errDir := os.Getwd(); errDir == nil {
		path = currentDir + "/" + projectName
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			fmt.Println("This project already exist")
			os.Exit(0)
		}
	}

	if e := template.CreateQtProjectTemplate(projectName); e != nil {
		fmt.Errorf("Operation is failed : %v", e)
	}

	out := os.Stdout
	if len(path) != 0 {
		err := tree.DirTree(out, path, true)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("PROJECT GENERATION COMPLETED")
}
