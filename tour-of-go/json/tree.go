package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type report struct {
	dirCount, fileCount int
}

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	var r report
	var res string
	var err error
	for _, arg := range args {
		r = report{}
		res, err = tree(arg, "", "", "", &r)
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
	res = fmt.Sprintf("%v\n%v directories, %v files", res, r.dirCount-(r.fileCount+1), r.fileCount)
	fmt.Println(res)
}

func tree(root, indent, line, res string, r *report) (string, error) {
	fi, err := os.Stat(root)
	if err != nil {
		return "", fmt.Errorf("could not stat %s: %v", root, err)
	}

	// fmt.Println(line + fi.Name())
	res += line + fi.Mode().Perm().String() + root + "/" + fi.Name() + "\n"
	r.dirCount++
	if !fi.IsDir() {
		r.fileCount++
		return res, nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return res, fmt.Errorf("could not read dir %s: %v", root, err)
	}

	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			// fmt.Printf(indent + "└──")
			line = indent + "└──"
			add = "   "
		} else {
			// fmt.Printf(indent + "├──")
			line = indent + "├──"
		}
		// res += line + " " + name + "\n"

		if res, err = tree(filepath.Join(root, name), indent+add, line, res, r); err != nil {
			return res, err
		}
	}

	return res, nil
}
