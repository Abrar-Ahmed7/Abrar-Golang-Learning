package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirTree(t *testing.T) {

	// creating dir
	if err := os.MkdirAll("resources/test-folder1/test-folder2", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	_, err := os.Create("resources/test-folder1/test-folder2/sample.txt")
	if err != nil {
		fmt.Printf("Unable to open file: %s", err)
	}
	_, err1 := os.Create("resources/sample1.txt")
	if err1 != nil {
		fmt.Printf("Unable to open file: %s", err)
	}

	dirPath := "./resources/"

	c := configs{}
	gotTree := getDirTree(dirPath, c)

	wantedTree := []string{
		"resources",
		"│\t│──sample1.txt",
		"│\t│──test-folder1",
		"│\t│\t│──test-folder2",
		"│\t│\t│\t│──sample.txt",
	}
	for i:=0; i<len(wantedTree); i++ {
		if gotTree[i] != wantedTree[i] {
			t.Errorf("Trees don't match")
		}
	}

	c = configs {
		dirOnly: true
	}
	gotDirTree := getDirTree(dirPath, c)

	wantedDirTree := []string{
		"resources",
		"│\t│──test-folder1",
		"│\t│\t│──test-folder2",
	}

	for i:=0; i<len(wantedDirTree); i++ {
		if gotDirTree[i] != wantedDirTree[i] {
			t.Errorf("Trees don't match")
		}
	}

	c = configs{
		perm: true
	}
	gotPermTree := getDirTree(dirPath, c)

	wantedPermTree := []string{
		"[-rwxrwxr-x]resources",
		"│\t│──[-rw-rw-r--]sample1.txt",
		"│\t│──[-rwxrwxr-x]test-folder1",
		"│\t│\t│──[-rwxrwxr-x]test-folder2",
		"│\t│\t│\t│──[-rw-rw-r--]sample.txt",
	}
	for i:=0; i<len(wantedPermTree); i++ {
		if gotPermTree[i] != wantedPermTree[i] {
			t.Errorf("Trees don't match")
		}
	}

	c = configs{
		level: 2
	}
	gotLevel2Tree := getDirTree(dirPath, c)

	wantedLevel2Tree := []string{
		"resources",
		"│\t│──sample1.txt",
		"│\t│──test-folder1",
	}
	for i:=0; i<len(wantedLevel2Tree); i++ {
		if gotLevel2Tree[i] != wantedLevel2Tree[i] {
			t.Errorf("Trees don't match")
		}
	}

	// deleting the created dir
	os.RemoveAll("./resources/")

}
