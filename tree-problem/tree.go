package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type report struct {
	dirCount, fileCount int
}

type configs struct {
	dirOnly, relPath, perm, json, xml bool
	level                             int
}

const (
	VER_LINE   = "│"
	HORI_LINE  = "──"
	VH_LINE    = VER_LINE + HORI_LINE
	END_LINE   = "└──"
	JSON_DIR   = "{\"type\":\"directory\",\"name\":\""
	JSON_FILE  = "{\"type\":\"file\",\"name\":\""
	XML_HEADER = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	XML_DIR    = "<directory name=\""
	XML_FILE   = "<file name=\""
	IN_OCT     = true
	IN_STR     = false
)

func main() {
	dirPath := os.Args[len(os.Args)-1]
	c := parseArgs()
	var res string
	var err error
	r := report{}

	if c.json {
		res, err = jsonTree(dirPath, "  ", "", &r, c, 1)
		if err != nil {
			log.Printf("tree %s: %v\n", dirPath, err)
		}

		fmt.Println(addSqBrkt(appendReport(res, r, c)))
	} else if c.xml {
		res, err = xmlTree(dirPath, "  ", "", &r, c, 1)
		if err != nil {
			log.Printf("tree %s: %v\n", dirPath, err)
		}

		fmt.Println(addTreeTag(appendReport(res, r, c)))
	} else {
		res, err = tree(dirPath, "", "", "", &r, c, 0)
		if err != nil {
			log.Printf("tree %s: %v\n", dirPath, err)
		}

		fmt.Println(appendReport(res, r, c))
	}
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
		if os.Args[i] == "-X" {
			c.xml = true
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
		add := VER_LINE + "  "
		if i == len(names)-1 {
			line = indent + END_LINE
			add = "   "
		} else {
			line = indent + VH_LINE
		}
		if res, err = tree(filepath.Join(root, name), indent+add, line, res, r, c, depth+1); err != nil {
			return res, err
		}
	}

	return res, nil
}

func jsonTree(root, indent, res string, r *report, c configs, depth int) (string, error) {
	fi, err := os.Stat(root)
	if err != nil {
		return "", fmt.Errorf("could not stat %s: %v", root, err)
	}

	res += indent + applyConfigsForJson(fi, root, c)
	r.dirCount++

	if !fi.IsDir() {
		r.fileCount++
		res += "\n"
		return res, nil
	}

	if c.level != 0 && c.level == depth-1 {
		return res, nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return res, fmt.Errorf("could not read dir %s: %v", root, err)
	}

	if len(fis) == 0 {
		res += "}\n"
	} else {
		res += ",\"contents\":[\n"
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
		if res, err = jsonTree(filepath.Join(root, name), strings.Repeat("  ", depth+1), res, r, c, depth+1); err != nil {
			return res, err
		}
		if i == len(names)-1 {
			res += indent + "]}\n"
		} else {
			l := len(res) - 1
			res = res[:l] + "," + res[l:]
		}
	}

	return res, nil
}

func xmlTree(root, indent, res string, r *report, c configs, depth int) (string, error) {
	fi, err := os.Stat(root)
	if err != nil {
		return "", fmt.Errorf("could not stat %s: %v", root, err)
	}

	res += indent + applyConfigsForXml(fi, root, c)
	r.dirCount++

	if !fi.IsDir() {
		r.fileCount++
		res += "\n"
		return res, nil
	}

	if c.level != 0 && c.level == depth-1 {
		return res, nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return res, fmt.Errorf("could not read dir %s: %v", root, err)
	}

	if len(fis) == 0 {
		res += "</directory>\n"
	} else {
		res += "\n"
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
		if res, err = xmlTree(filepath.Join(root, name), strings.Repeat("  ", depth+1), res, r, c, depth+1); err != nil {

			return res, err
		}
		if i == len(names)-1 {
			res += indent + "</directory>\n"
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

func applyConfigsForJson(file os.FileInfo, path string, c configs) string {
	if c.relPath && c.perm {
		if file.IsDir() {
			return JSON_DIR + path + "/" + file.Name() + ",\"mode\":\"" + getMode(file, IN_OCT) + "\",\"prot\":\"" + getMode(file, IN_STR) + "\""
		}
		return JSON_FILE + path + "/" + file.Name() + ",\"mode\":\"" + getMode(file, IN_OCT) + "\",\"prot\":\"" + getMode(file, IN_STR) + "\"}"
	} else if !c.relPath && c.perm {
		if file.IsDir() {
			return JSON_DIR + file.Name() + ",\"mode\":\"" + getMode(file, IN_OCT) + "\",\"prot\":\"" + getMode(file, IN_STR) + "\""
		}
		return JSON_FILE + file.Name() + ",\"mode\":\"" + getMode(file, IN_OCT) + "\",\"prot\":\"" + getMode(file, IN_STR) + "\"}"
	} else if c.relPath && !c.perm {
		if file.IsDir() {
			return JSON_DIR + path + "/" + file.Name() + "\""
		}
		return JSON_FILE + path + "/" + file.Name() + "\"}"
	} else {
		if file.IsDir() {
			return JSON_DIR + file.Name() + "\""
		}
		return JSON_FILE + file.Name() + "\"}"
	}
}

func applyConfigsForXml(file os.FileInfo, path string, c configs) string {
	if c.relPath && c.perm {
		if file.IsDir() {
			return XML_DIR + path + "/" + file.Name() + ",mode=\"" + getMode(file, IN_OCT) + "\",prot=\"" + getMode(file, IN_STR) + "\">"
		}
		return XML_FILE + path + "/" + file.Name() + ",mode=\"" + getMode(file, IN_OCT) + "\",prot=\"" + getMode(file, IN_STR) + "\"></file>"
	} else if !c.relPath && c.perm {
		if file.IsDir() {
			return XML_DIR + file.Name() + ",mode=\"" + getMode(file, IN_OCT) + "\",prot=\"" + getMode(file, IN_STR) + "\">"
		}
		return XML_FILE + file.Name() + ",mode=\"" + getMode(file, IN_OCT) + "\",prot=\"" + getMode(file, IN_STR) + "\"></file>"
	} else if c.relPath && !c.perm {
		if file.IsDir() {
			return XML_DIR + path + "/" + file.Name() + "\">"
		}
		return XML_FILE + path + "/" + file.Name() + "\"></file>"
	} else {
		if file.IsDir() {
			return XML_DIR + file.Name() + "\">"
		}
		return XML_FILE + file.Name() + "\"></file>"
	}
}

func appendReport(res string, r report, c configs) string {
	if c.json {
		res += ",\n"
		if c.dirOnly {
			res = fmt.Sprintf("  {\"type\":\"report\",\"directories\":%v}", r.dirCount-(r.fileCount+1))
		} else {
			res += fmt.Sprintf("  {\"type\":\"report\",\"directories\":%v, \"files\":%v}", r.dirCount-(r.fileCount+1), r.fileCount)
		}
	} else if c.xml {
		res += "  <report>\n    <directories>"
		if c.dirOnly {
			res += fmt.Sprintf("%v</directories>", r.dirCount-(r.fileCount+1))
		} else {
			res += fmt.Sprintf("%v</directories>\n    <files>%v</files>", r.dirCount-(r.fileCount+1), r.fileCount)
		}
		res += "\n  </report>"
	} else {
		if c.dirOnly {
			res = fmt.Sprintf("%v\n%v directories", res, r.dirCount-(r.fileCount+1))
		} else {
			res = fmt.Sprintf("%v\n%v directories, %v files", res, r.dirCount-(r.fileCount+1), r.fileCount)
		}
	}
	return res
}

func getMode(file os.FileInfo, inOct bool) string {
	if inOct {
		return fmt.Sprintf("%#o", file.Mode().Perm())
	}

	return file.Mode().Perm().String()
}

func addSqBrkt(res string) string {
	res = "[\n" + res + "\n]"
	return res
}

func addTreeTag(res string) string {
	res = XML_HEADER + "\n<tree>\n" + res + "\n</tree>"
	return res
}

func convertStrToInt(str string) int {
	var num, err = strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
