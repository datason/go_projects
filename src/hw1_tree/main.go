package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// func recursiveDirTree()

func dirTree(output io.Writer, path string, printFiles bool, tabCounter int) error {
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

	for e, file := range fileInfo {

		if file.IsDir() {
			startStr := strings.Repeat("|   ", tabCounter)

			if e < (len(fileInfo) - 1) {
				res := startStr + "├───" + file.Name() + "\n"
				fmt.Printf(res)

			} else {
				res := startStr + "└───" + file.Name() + "\n"
				fmt.Printf(res)
			}
			dirTree(output, path+`/`+file.Name(), printFiles, tabCounter+1)
		} else if printFiles {
			startStr := strings.Repeat("|   ", tabCounter)
			// print("%T", file.Size())
			var endStr string
			if file.Size() == 0 {
				endStr = " (empty)"
			} else {
				endStr = fmt.Sprintf(" (%db)", file.Size()) // (10b)
			}

			if e < (len(fileInfo) - 1) {
				res := startStr + "├───" + file.Name() + endStr + "\n"
				fmt.Printf(res)
			} else {
				res := startStr + "└───" + file.Name() + endStr + "\n"
				fmt.Printf(res)
			}
		} else {
			continue
		}

	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles, 0)
	if err != nil {
		panic(err.Error())
	}
}
