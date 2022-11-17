package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type commands struct {
	d, f, p string
}

func main() {
	dirPath := "./"
	c := commands{
		d: "",
		f: "",
		p: "",
	}

	printDirFiles(dirPath, c)
}

func printDirFiles(dirPath string, c commands) {
	var dirCount, fileCount int
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			for i := 0; i < len(currentPath); i++ {
				if path == dirPath {
					continue
				}
				tabSpace = "|--" + tabSpace
			}
			if file.IsDir() {
				dirCount++
				if c.d == "d" {
					applyCommands(tabSpace, file, path, c)
				}
			}
			fileCount++
			if c.d != "d" {
				applyCommands(tabSpace, file, path, c)
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	if c.d == "d" {
		fmt.Println("\n", dirCount, "directories")
	} else {
		fmt.Println("\n", dirCount, "directories,", fileCount-dirCount, "files")
	}
}

func applyCommands(tabSpace string, file os.FileInfo, path string, c commands) {
	if c.f == "f" && c.p == "p" {
		fmt.Printf("%v [%v] %v\n ", tabSpace, file.Mode().Perm(), path)
	} else if c.f != "f" && c.p == "p" {
		fmt.Printf("%v [%v] %v\n ", tabSpace, file.Mode().Perm(), file.Name())
	} else if c.f == "f" && c.p != "p" {
		fmt.Println(tabSpace, path)
	} else {
		fmt.Println(tabSpace, file.Name())
	}
}
