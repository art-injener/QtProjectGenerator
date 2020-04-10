package tree

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Node interface {
	fmt.Stringer
}

type File struct {
	name string
	size int64
}

func (file File) String() string {
	if file.size == 0 {
		return file.name + " (empty)"
	}
	return file.name + " (" + strconv.FormatInt(file.size, 10) + "b)"
}

type Dir struct {
	name     string
	children []Node
}

func (dir Dir) String() string {
	return dir.name
}

func DirTree(out io.Writer, path string, files bool) error {

	nodes, err := scanDir(path, []Node{}, files)
	printDir(out, nodes, []string{})

	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func scanDir(path string, nodes []Node, withFiles bool) ([]Node, error) {
	if len(path) == 0 {
		return nil, errors.New("Empty filepath")
	}
	file, err := os.Open(path)
	if file == nil || err != nil {
		return nil, errors.New("Error open file")
	}

	files, err := file.Readdir(0)
	file.Close()

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, info := range files {
		if !(info.IsDir() || withFiles) {
			continue
		}
		var newNode Node

		if info.IsDir() {
			children, _ := scanDir(filepath.Join(path, info.Name()), []Node{}, withFiles)
			newNode = Dir{info.Name(), children}
		} else {
			newNode = File{info.Name(), info.Size()}
		}
		nodes = append(nodes, newNode)
	}

	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func printDir(out io.Writer, nodes []Node, prefixes []string) {
	if len(nodes) == 0 {
		return
	}

	fmt.Fprintf(out, "%s", strings.Join(prefixes, ""))

	node := nodes[0]

	if len(nodes) == 1 {
		fmt.Fprintf(out, "%s%s\n", "└───", node)
		if directory, ok := node.(Dir); ok {
			printDir(out, directory.children, append(prefixes, "\t"))
		}
		return
	}

	fmt.Fprintf(out, "%s%s\n", "├───", node)
	if directory, ok := node.(Dir); ok {
		printDir(out, directory.children, append(prefixes, "│\t"))
	}

	printDir(out, nodes[1:], prefixes)
}
