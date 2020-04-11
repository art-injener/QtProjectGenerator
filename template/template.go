package template

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type templateDirTree struct {
	nameRootDir string
	currentDir  string
	listDirs    []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateQtProjectTemplate(name string) error {
	if len(name) == 0 {
		return errors.New("The string with the project name is empty")
	}

	textFile := []byte("File for add dir to git\n")
	if currentDir, errDir := os.Getwd(); errDir == nil {
		if strdir, err := newTemplateDirTree(name, currentDir); err == nil {
			//defer os.RemoveAll(currentDir+"/"+name)
			fmt.Println(strdir.nameRootDir)
			for _, v := range strdir.listDirs {
				var errMkDir = os.MkdirAll(v, 0766)
				check(errMkDir)

				err := ioutil.WriteFile(v+"/README.txt", textFile, 0644)
				check(err)
			}
			mapDirsData := getMapFiles(strdir.nameRootDir, name)
			for k, v := range mapDirsData {
				err := ioutil.WriteFile(k, v, 0644)
				check(err)
			}
		} else {
			return err
		}
	} else {
		return errDir
	}

	return nil
}

func newTemplateDirTree(name string, curDir string) (*templateDirTree, error) {
	if len(name) == 0 {
		return nil, errors.New("The string with the project name is empty")
	}
	p := templateDirTree{nameRootDir: name,
		currentDir: curDir,
		listDirs:   getSliceDirs(curDir + "/" + name)}
	return &p, nil
}

func getSliceDirs(path string) []string {
	return []string{path + "/bin",
		path + "/build",
		path + "/config",
		path + "/data",
		path + "/doc",
		path + "/export",
		path + "/import",
		path + "/lib.linux",
		path + "/lib.win",
		path + "/src",
		path + "/tests",
		path + "/utils",
		path + "/src/Apps",
		path + "/src/Libs",
		path + "/src/include"}
}

func getMapFiles(path, projectName string) map[string][]byte {

	m := make(map[string][]byte)
	m[path+"/"+projectName+".pro"] = rootProFile
	m[path+"/common.pri"] = commonPriFile
	m[path+"/lib.pri"] = libPriFile
	m[path+"/app.pri"] = appPriFile
	m[path+"/src/Apps/Apps.pro"] = subAppProFile
	m[path+"/src/Libs/Libs.pro"] = subLibProFile
	return m
}
