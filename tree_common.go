package tree

import (
	"os"
	"strings"
	"time"
)

func getFilepath(name, path string) string {
	filename := getFilename(name)
	filepath := checkingAndCreatingDirectory(path)
	return filepath + filename
}

func getFilename(name string) string {
	// Checking for file extension
	var filename string
	var nameFragments = strings.Split(name, ".")
	switch {
	case len(name) == 0:
		filename = "tree_" + time.Now().Format("01-01-2006") + ".json"
	case len(nameFragments) > 1:
		filename = nameFragments[0] + ".json"
	default:
		filename = name + ".json"
	}
	return filename
}

func checkingAndCreatingDirectory(path string) string {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		os.Mkdir("./trees/", 0777)
		return "./trees/"
	}
	if path[len(path)-1] != '/' {
		return path + "/"
	}
	return path
}
