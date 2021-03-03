package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func recursiveDirTree(output io.Writer, path string, printFiles bool, startStr string) error {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(fileInfo, func(i, j int) bool {
		return fileInfo[i].Name() < fileInfo[j].Name()
	})

	// Filter before main print loop, because of true last position.
	// If we don't print files, but they are in os.FileInfo, last folder can be print uncorrectly.
	if !printFiles {
		var fileInfoFiltered []os.FileInfo

		for _, file := range fileInfo {
			if file.IsDir() {
				fileInfoFiltered = append(fileInfoFiltered, file)
			}
		}
		fileInfo = fileInfoFiltered
	}

	for e, file := range fileInfo {

		if file.IsDir() {

			if e < (len(fileInfo) - 1) {
				res := startStr + "├───" + file.Name() + "\n"
				fmt.Fprintf(output, res)
				recursiveDirTree(output, path+`/`+file.Name(), printFiles, startStr+"│	")
			} else {
				res := startStr + "└───" + file.Name() + "\n"
				fmt.Fprintf(output, res)

				recursiveDirTree(output, path+`/`+file.Name(), printFiles, startStr+"	")
			}

		} else if printFiles {
			var endStr string
			if file.Size() == 0 {
				endStr = " (empty)"
			} else {
				endStr = fmt.Sprintf(" (%db)", file.Size()) // (10b)
			}

			if e < (len(fileInfo) - 1) {
				res := startStr + "├───" + file.Name() + endStr + "\n"
				fmt.Fprintf(output, res)
			} else {
				res := startStr + "└───" + file.Name() + endStr + "\n"
				fmt.Fprintf(output, res)
			}
		} else {
			continue
		}

	}
	return nil
}

func dirTree(output io.Writer, path string, printFiles bool) error {
	return recursiveDirTree(output, path, printFiles, "")
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
