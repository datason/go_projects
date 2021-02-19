package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// func walkFn(path string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
// 		return err
// 	}
// 	if info.IsDir() {
// 		fmt.Printf("visited dir: %q\n", path)
// 	} else {
// 		fmt.Printf("visited file: %q\n", path)
// 	}
// 	return nil
// }

// func dirTree(output io.Writer, path string, printFiles bool) error {
// 	err = filepath.Walk(".", walkFn(path))

// 	if err != nil {
// 		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
// 		return
// 	}
// }

// func recursiveDirTree()

func dirTree(output io.Writer, path string, printFiles bool, tab_counter int) error {
	// print(fmt.Println(reflect.TypeOf(output)))
	// print(fmt.Println(reflect.TypeOf(path)))
	// print(fmt.Println(reflect.TypeOf(printFiles)))

	// printFiles: true: Выводит дерево каталогов и файлов
	// output: Сюда надо вывести всю писанину.
	// path: Для какой директории это надо сделать
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Current dir is", dir)

	// print(fmt.Println(reflect.TypeOf(f)))

	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(fileInfo, func(i, j int) bool {
		return fileInfo[i].Name() < fileInfo[j].Name()
	})

	var files []string
	for e, file := range fileInfo {
		// print(fmt.Println(reflect.TypeOf(file)))
		// fmt.Printf("%d", e)
		// if e == len(fileInfo) -1

		if file.IsDir() {
			// fmt.Printf("It is dir %q \n", file.Name())
			startStr := strings.Repeat("|   ", tab_counter)
			// fmt.Printf("Tab counter %d", tab_counter)

			if e < (len(fileInfo) - 1) {
				// print("%d %d", e, len(fileInfo))
				res := startStr + "├───" + file.Name() + "\n"
				fmt.Printf(res)

				// } else if e == 0 {
				// 	res := start_str + "└───" + file.Name() + "\n"
				// 	fmt.Printf(res)
			} else {
				res := startStr + "└───" + file.Name() + "\n"
				fmt.Printf(res)
			}
			dirTree(output, path+`/`+file.Name(), printFiles, tab_counter+1)
		} else {

			startStr := strings.Repeat("|   ", tab_counter)
			// fmt.Printf("Tab counter %d", tab_counter)

			if e < (len(fileInfo) - 1) {
				// print("%d %d", e, len(fileInfo))
				res := startStr + "├───" + file.Name() + "\n"
				fmt.Printf(res)

				// } else if e == 0 {
				// 	res := start_str + "└───" + file.Name() + "\n"
				// 	fmt.Printf(res)
			} else {
				res := startStr + "└───" + file.Name() + "\n"
				fmt.Printf(res)
			}
		}

		// fmt.Printf("\n")
		files = append(files, file.Name())

	}
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

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
