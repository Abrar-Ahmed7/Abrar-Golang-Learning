package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dirPath := "./"
	fmt.Println("---------------Directories amd Files---------------")
	printDirAndFiles(dirPath)
	fmt.Println("---------------Directories---------------")
	printDirAlone(dirPath)
	fmt.Println("---------------Directories and Files with Realtive Path---------------")
	printDirAndFilesWithRelPath(dirPath)
	fmt.Println("---------------Directories Alone with Relative Path---------------")
	printDirAloneWithRelPath(dirPath)
	fmt.Println("---------------Directories and Files with Permission---------------")
	printDirAndFilesWithPerm(dirPath)
}

func printDirAndFiles(dirPath string) {
	var dirCount, fileCount int
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			for i := 0; i < len(currentPath); i++ {
				if path == "./" {
					continue
				}
				tabSpace = "|--" + tabSpace
			}
			if file.IsDir() {
				dirCount++
			}
			fileCount++
			fmt.Println(tabSpace, file.Name())
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n", dirCount, "directories,", fileCount, "files")
}

func printDirAlone(dirPath string) {
	var dirCount int
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			for i := 0; i < len(currentPath); i++ {
				if path == "./" {
					continue
				}
				tabSpace = "|--" + tabSpace
			}
			if file.IsDir() {
				dirCount++
				fmt.Println(tabSpace, file.Name())
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n", dirCount, "directories")
}

func printDirAndFilesWithRelPath(dirPath string) {
	var dirCount, fileCount int
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			for i := 0; i < len(currentPath); i++ {
				if path == "./" {
					continue
				}
				tabSpace = "|--" + tabSpace
			}
			if file.IsDir() {
				dirCount++
			}
			fileCount++
			fmt.Println(tabSpace, path)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n", dirCount, "directories,", fileCount, "files")
}

func printDirAloneWithRelPath(dirPath string) {
	var dirCount int
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			for i := 0; i < len(currentPath); i++ {
				if path == "./" {
					continue
				}
				tabSpace = "|--" + tabSpace
			}
			if file.IsDir() {
				dirCount++
				fmt.Println(tabSpace, path)
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n", dirCount, "directories")
}

func printDirAndFilesWithPerm(dirPath string) {
	var dirCount, fileCount int
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			for i := 0; i < len(currentPath); i++ {
				if path == "./" {
					continue
				}
				tabSpace = "|--" + tabSpace
			}
			if file.IsDir() {
				dirCount++
			}
			fileCount++
			fmt.Printf("\n%v [%v] %v", tabSpace, file.Mode().Perm(), file.Name())
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n", dirCount, "directories,", fileCount, "files")
}
