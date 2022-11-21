package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"container/list"
)

type configs struct {
	dirOnly, relPath, perm bool
	level                  int
}

const (
	verLine = "│"
	horLine = "──"
	vhLine  = verLine + horLine
)

func main() {
	dirPath := os.Args[len(os.Args)-1]
	c := parseArgs()
	printDirTree(dirPath, c)
}

func parseArgs() configs {
	c := configs{}
	for i := 0; i < len(os.Args)-1; i++ {
		if os.Args[i] == "-d" {
			c.dirOnly = true
		}
		if os.Args[i] == "-f" {
			c.relPath = true
		}
		if os.Args[i] == "-p" {
			c.perm = true
		}
		if os.Args[i] == "-L" {
			c.level = convertStrToInt(os.Args[i+1])
		}
	}
	return c
}

func printDirTree(dirPath string, c configs) {
	dirTree := getDirTree(dirPath, c)
	for _, d := range dirTree {
		fmt.Println(d)
	}
}

func getDirTree(dirPath string, c configs) []string {
	var dirCount, fileCount int
	var dirTree []string
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			var level int
			if c.level != 0 {
				level = c.level
			} else {
				level = len(currentPath)
			}
			for i := 0; i < len(currentPath); i++ {
				if path == dirPath {
					continue
				}
				if i == len(currentPath)-1 {
					tabSpace = tabSpace + vhLine
				} else {
					tabSpace = verLine + "\t" + tabSpace
				}

			}
			if file.IsDir() {
				dirCount++
				if c.dirOnly {
					if len(currentPath) <= level {
						dirTree = append(dirTree, applyConfigs(tabSpace, file, path, c))
						// fmt.Println(applyCommands(tabSpace, file, path, c))
					}
				}
			}
			fileCount++
			if !c.dirOnly {
				if len(currentPath) <= level {
					dirTree = append(dirTree, applyConfigs(tabSpace, file, path, c))
					// fmt.Println(applyConfigs(tabSpace, file, path, c))
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	if c.dirOnly {
		dirTree = append(dirTree, "\n"+strconv.Itoa(dirCount)+" directories")
		// fmt.Println("\n", dirCount, "directories")
	} else {
		dirTree = append(dirTree, "\n"+strconv.Itoa(dirCount)+" directories, "+strconv.Itoa(fileCount-dirCount)+" files")
	}
	return dirTree
}

func applyConfigs(tabSpace string, file os.FileInfo, path string, c configs) string {
	if c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + path
	} else if !c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + file.Name()
	} else if c.relPath && !c.perm {
		return tabSpace + path
	} else {
		return tabSpace + file.Name()
	}
}

func convertStrToInt(str string) int {
	var num, err = strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

/*
My notes:
	for cli:
		- take all the arguments or flags from cli
		- assign it to the respective variable inside the struct
		- the unassigned variable will be initialized with nil or "" or 0
		- if they are not initialized assign them with nil or "" or 0 inside the struct

	for adding levels:
		- Have to find where to stop printing
		- or should I go for a new approach?
*/

/*
	- Below lines of code are just to try the things bit differently
	- And, trying to implement other stories as well
*/

func approach2(path string) {
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf(home)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			path = path + "/" + f.Name()
			approach2(path)
		}
		fmt.Println(f.Name())
	}
}

func printListing(entry string, depth int) {
	indent := strings.Repeat("|   ", depth)
	fmt.Printf("%s|-- %s\n", indent, entry)
}

func printDirectory(path string, depth int) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	printListing(path, depth)
	for _, entry := range entries {
		if (entry.Mode() & os.ModeSymlink) == os.ModeSymlink {
			full_path, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Printf("error reading link: %s\n", err.Error())
			} else {
				printListing(entry.Name()+" -> "+full_path, depth+1)
			}
		} else if entry.IsDir() {
			printDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {
			printListing(entry.Name(), depth+1)
		}
	}
}

func printDir(path string) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", "path", err.Error())
	}
	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func printJson(dirPath string, c configs) {
	var dirCount, fileCount int
	var jsonTree []string
	jsonTree = append(jsonTree, "[")
	err := filepath.Walk(dirPath,
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var tabSpace string
			currentPath := strings.Split(path, "/")
			var level int
			if c.level != 0 {
				level = c.level
			} else {
				level = len(currentPath)
			}
			for i := 0; i < len(currentPath); i++ {
				if path == dirPath {
					continue
				}
				if i == len(currentPath)-1 {
					// fmt.Println(i)
					tabSpace = tabSpace + "\t"
				} else {
					tabSpace = "\t" + tabSpace
				}

			}
			if file.IsDir() {
				dirCount++
				if c.dirOnly {
					if len(currentPath) <= level {
						jsonTree = append(jsonTree, applyCommands1(tabSpace, file, path, c))
						// fmt.Println(applyCommands(tabSpace, file, path, c))
					}
				}
			}
			fileCount++
			if !c.dirOnly {
				if len(currentPath) <= level {
					jsonTree = append(jsonTree, applyCommands1(tabSpace, file, path, c))
				}
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	if c.dirOnly {
		fmt.Println("\n", dirCount, "directories")
	} else {
		fmt.Println("\n", dirCount, "directories,", fileCount-dirCount, "files")
	}
}

func applyCommands1(tabSpace string, file os.FileInfo, path string, c configs) string {
	if c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + path
	} else if !c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + file.Name()
	} else if c.relPath && !c.perm {
		return tabSpace + path
	} else {
		return tabSpace + file.Name()
	}
}

type Node struct {
	Data                string
	parent              *Node
	children            []*Node
	dirCount, fileCount int
}

func (n Node) initiate(path string) Node {
	n.Data = path
	n.children = []*Node{}
	n.dirCount = 0
	n.fileCount = 0
	return n
}

func rootDirTree(folder string, c configs) {
	n := Node{
		// Data: "",
		// parent: &Node{},
		// children: []*Node{},
		// dirCount: 0,
		// fileCount: 0,
	}
	dirRoot := n.initiate(folder)
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", "path", err.Error())
	}
	for _, f := range files {
		if f.IsDir() {
			dirRoot.dirCount++
			appendDirTree(f.Name(), dirRoot, c)
		}
	}

}

func appendDirTree(file string, dirRoot Node, c configs) {
	dirRoot.addChild(file)
	files, err := ioutil.ReadDir(file)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", "path", err.Error())
	}
	for _, f := range files {
		if f.IsDir() {
			dirRoot.dirCount++
			appendDirTree(f.Name(), *dirRoot.children[len(dirRoot.children)-1], c)
		} else {
			dirRoot.fileCount++
			appendFile(f.Name(), *dirRoot.children[len(dirRoot.children)-1], c)
		}
	}
}

func appendFile(file string, fileNode Node, c configs) {
	if c.dirOnly {
		return
	}
	fileNode.addChild(file)
}

func (n Node) addChild(folder string) {
	childNode := Node{}
	childNode.initiate(folder)
	childNode.parent = &n
	childNode.children = append(childNode.children, &childNode)
}

func getDirectoryTree(tree Node, c configs) list.List {
	result := list.New()
	// if c.relPath && c.perm {
	// 	addRelPathAndPerm(result, tree.Data)
	// }
	result.PushFront(tree.Data)
	for i := 0; i < len(tree.children); i++ {
		if hasNext(tree.children, tree.children[i]) {
			subTree := getDirectoryTree(*tree.children[i+1], c)
			addSubTree(result, subTree)
		}
	}
	return *result
}

func addSubTree(result *list.List, subTree list.List) {
	// result.InsertAfter()
}

func hasNext(children []*Node, child *Node) bool {
	for i, c := range children {
		if c == child && children[i+1] != nil {
			return true
		}
	}
	return false
}

func next(children []*Node, child *Node) bool {
	for i, c := range children {
		if c == child && children[i+1] != nil {
			return true
		}
	}
	return false
}

func addRelPathAndPerm(result *list.List, s string) {

}

// type UserCollection struct {
//     users []*User
// }

// func (u *Node) createIterator() Iterator {
//     return &Node.children{
//     }
// }

// type Iterator interface {
//     hasNext() bool
//     getNext() *User
// }

// type nodeIterator struct {
// 	index int
// 	nodes []*Node
// }

// type UserIterator struct {
//     index int
//     users []*User
// }

// func (u *UserIterator) hasNext() bool {
//     if u.index < len(u.users) {
//         return true
//     }
//     return false

// }
// func (u *UserIterator) getNext() *User {
//     if u.hasNext() {
//         user := u.users[u.index]
//         u.index++
//         return user
//     }
//     return nil
// }
