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

type report struct {
	dirCount, fileCount int
}

type configs struct {
	dirOnly, relPath, perm, json bool
	level                        int
}

//	type struct {
//		root, indent, line, res string
//	}
const (
	verLine = "│"
	horLine = "──"
	vhLine  = verLine + horLine
	endLine = "└──"
)

func main() {
	dirPath := os.Args[len(os.Args)-1]
	c := parseArgs()
	// printTree(dirPath, c)
	// var r report
	// var res string
	// var err error
	r := report{}
	res, err := tree(dirPath, "", "", "", &r, c, 0)
	if err != nil {
		log.Printf("tree %s: %v\n", dirPath, err)
	}
	fmt.Println(appendReport(res, r, c))
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
		if os.Args[i] == "-J" {
			c.json = true
		}
	}
	return c
}

func tree(root, indent, line, res string, r *report, c configs, depth int) (string, error) {
	fi, err := os.Stat(root)
	if err != nil {
		return "", fmt.Errorf("could not stat %s: %v", root, err)
	}
	res += line + applyConfigs(fi, root, c) + "\n"
	r.dirCount++
	if !fi.IsDir() {
		r.fileCount++
		return res, nil
	}

	if c.level != 0 && c.level == depth {
		return res, nil
	}
	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return res, fmt.Errorf("could not read dir %s: %v", root, err)
	}
	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			if c.dirOnly && !fi.IsDir() {
				continue
			}
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		add := verLine + "  "
		if i == len(names)-1 {
			line = indent + endLine
			add = "   "
		} else {
			line = indent + vhLine
		}
		if res, err = tree(filepath.Join(root, name), indent+add, line, res, r, c, depth+1); err != nil {
			return res, err
		}
	}
	return res, nil
}

func applyConfigs(file os.FileInfo, path string, c configs) string {
	if c.relPath && c.perm {
		return "[" + file.Mode().Perm().String() + "]" + path + "/" + file.Name()
	} else if !c.relPath && c.perm {
		return "[" + file.Mode().Perm().String() + "]" + file.Name()
	} else if c.relPath && !c.perm {
		return path + "/" + file.Name()
	} else {
		return file.Name()
	}
}

func appendReport(res string, r report, c configs) string {
	if c.dirOnly {
		res = fmt.Sprintf("%v\n%v directories", res, r.dirCount-(r.fileCount+1))
	} else {
		res = fmt.Sprintf("%v\n%v directories, %v files", res, r.dirCount-(r.fileCount+1), r.fileCount)
	}
	return res
}

// the new approach stops here

func printTree(dirPath string, c configs) {
	var dirTree []string
	if c.json == true {
		dirTree = getJsonTree(dirPath, c)
	} else {
		dirTree = getDirTree(dirPath, c)
	}
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
						dirTree = append(dirTree, applyConfigs1(tabSpace, file, path, c))
					}
				}
			}
			fileCount++
			if !c.dirOnly {
				if len(currentPath) <= level {
					dirTree = append(dirTree, applyConfigs1(tabSpace, file, path, c))
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	if c.dirOnly {
		dirTree = append(dirTree, "\n"+strconv.Itoa(dirCount)+" directories")
	} else {
		dirTree = append(dirTree, "\n"+strconv.Itoa(dirCount)+" directories, "+strconv.Itoa(fileCount-dirCount)+" files")
	}
	return dirTree
}

func applyConfigs1(tabSpace string, file os.FileInfo, path string, c configs) string {
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

func getJsonTree(dirPath string, c configs) []string {
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
				tabSpace = tabSpace + "  "
			}
			if file.IsDir() {
				dirCount++
				if c.dirOnly {
					if len(currentPath) <= level {
						jsonTree = append(jsonTree, applyConfigsForJson(tabSpace+"{\"type\":\"directory\",\"name\":", file, path, c))
						// fmt.Println(applyCommands(tabSpace, file, path, c))
					}
				}
			}
			fileCount++
			if !c.dirOnly {
				if len(currentPath) <= level {
					if file.IsDir() {
						jsonTree = append(jsonTree, applyConfigsForJson(tabSpace+"{\"type\":\"directory\",\"name\":", file, path, c))
					} else {
						jsonTree = append(jsonTree, applyConfigsForJson(tabSpace+"{\"type\":\"file\",\"name\":", file, path, c))
					}
				}
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	if c.dirOnly {
		jsonTree = append(jsonTree, fmt.Sprintf("{\"type\":\"report\",\"directories\": %d}", dirCount))
	} else {
		jsonTree = append(jsonTree, fmt.Sprintf("{\"type\":\"report\",\"directories\": %d \"files\":%d}", dirCount, fileCount-dirCount))
	}
	jsonTree = append(jsonTree, "]")
	return jsonTree
}

func applyConfigsForJson(tabSpace string, file os.FileInfo, path string, c configs) string {
	if c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + path + "\"}"
	} else if !c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + file.Name() + "\"}"
	} else if c.relPath && !c.perm {
		return tabSpace + path + "\"}"
	} else {
		return tabSpace + "\"" + file.Name() + "\"}"
	}
}

func getXmlTree(dirPath string, c configs) []string {
	var dirCount, fileCount int
	var xmlTree []string
	xmlTree = append(xmlTree, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	xmlTree = append(xmlTree, "<tree>")
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
				tabSpace = tabSpace + "  "
			}
			if file.IsDir() {
				dirCount++
				if c.dirOnly {
					if len(currentPath) <= level {
						xmlTree = append(xmlTree, applyConfigsForJson(tabSpace+"{\"type\":\"directory\",\"name\":", file, path, c))
						// fmt.Println(applyCommands(tabSpace, file, path, c))
					}
				}
			}
			fileCount++
			if !c.dirOnly {
				if len(currentPath) <= level {
					if file.IsDir() {
						xmlTree = append(xmlTree, applyConfigsForJson(tabSpace+"{\"type\":\"directory\",\"name\":", file, path, c))
					} else {
						xmlTree = append(xmlTree, applyConfigsForJson(tabSpace+"{\"type\":\"file\",\"name\":", file, path, c))
					}
				}
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	if c.dirOnly {
		xmlTree = append(xmlTree, fmt.Sprintf("{\"type\":\"report\",\"directories\": %d}", dirCount))
	} else {
		xmlTree = append(xmlTree, fmt.Sprintf("{\"type\":\"report\",\"directories\": %d \"files\":%d}", dirCount, fileCount-dirCount))
	}
	xmlTree = append(xmlTree, "</tree>")
	return xmlTree
}

func applyConfigsForXml(tabSpace string, file os.FileInfo, path string, c configs) string {
	if c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + path + "\"}"
	} else if !c.relPath && c.perm {
		return tabSpace + "[" + file.Mode().Perm().String() + "]" + file.Name() + "\"}"
	} else if c.relPath && !c.perm {
		return tabSpace + path + "\"}"
	} else {
		return tabSpace + "\"" + file.Name() + "\"}"
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

func next(children []*Node, child *Node) *Node {
	for i, c := range children {
		if c == child && children[i+1] != nil {
			return children[i+1]
		}
	}
	return &Node{}
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
