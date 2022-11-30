package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"reflect"
// 	"testing"
// )

// func TestTree(t *testing.T) {

// 	// creating dir
// 	if err := os.MkdirAll("resources/test-folder1/test-folder2", os.ModePerm); err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err1 := os.Create("resources/test-folder1/test-folder2/sample.txt")
// 	if err1 != nil {
// 		fmt.Printf("Unable to open file: %s", err1)
// 	}
// 	_, err2 := os.Create("resources/sample1.txt")
// 	if err2 != nil {
// 		fmt.Printf("Unable to open file: %s", err2)
// 	}

//     tests := []struct {
// 		dirPath string
//         input configs
// 		r report
//         want  string
//     }{
//         {
// 			dirPath: "./resources",
// 		 	input: configs{},
// 			report: &r,
// 			want: "resources\n" +
//             "│──sample1.txt\n" +
//             "└──test-folder1\n" +
//             "   └──test-folder2\n" +
//             "      └──sample.txt\n" +
//             "\n" +
//             "2 directories, 2 files",
// 		},
//         {
// 			dirPath: "./resources",
// 		 	input: configs{dirOnly: true},
// 			report: &r,
// 			want: "resources\n" +
//             "└──test-folder1\n" +
//             "   └──test-folder2\n" +
//             "\n" +
//             "2 directories",
// 		},
//         {
// 			dirPath: "./resources",
// 		 	input: configs{level: 1},
// 			report: &r,
// 			want: "resources\n" +
//             "│──sample1.txt\n" +
//             "└──test-folder1\n" +
//             "\n" +
//             "1 directories, 1 files",
// 		},
//     }

//     for _, tc := range tests {
//         got := tree(tc.dirPath,"","","",&tc.r,tc.input,0)
//         if !reflect.DeepEqual(tc.want, got) {
//             t.Fatalf("expected: %v, got: %v", tc.want, got)
//         }
//     }

// 	// deleting the created dir
// 	os.RemoveAll("./resources/")
// }

// func TestGetDirTree(t *testing.T) {

// 	// creating dir
// 	if err := os.MkdirAll("resources/test-folder1/test-folder2", os.ModePerm); err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err := os.Create("resources/test-folder1/test-folder2/sample.txt")
// 	if err != nil {
// 		fmt.Printf("Unable to open file: %s", err)
// 	}
// 	_, err := os.Create("resources/sample1.txt")
// 	if err != nil {
// 		fmt.Printf("Unable to open file: %s", err)
// 	}

// 	dirPath := "./resources/"
// 	var got string
// 	var want string
// 	var c configs
// 	got = tree(dirPath,"", "", "", &r, c, 0)

// 	want =

// 	wantedTree := []string{
// 		"resources",
// 		"│\t│──sample1.txt",
// 		"│\t│──test-folder1",
// 		"│\t│\t│──test-folder2",
// 		"│\t│\t│\t│──sample.txt",
// 	}
// 	for i:=0; i<len(wantedTree); i++ {
// 		if gotTree[i] != wantedTree[i] {
// 			t.Errorf("Trees don't match")
// 		}
// 	}
// 	c.dirOnly= true
// 	c = configs {
// 		dirOnly: true,
// 	}
// 	gotDirTree := getDirTree(dirPath, c)

// 	wantedDirTree := []string{
// 		"resources",
// 		"│\t│──test-folder1",
// 		"│\t│\t│──test-folder2",
// 	}

// 	for i:=0; i<len(wantedDirTree); i++ {
// 		if gotDirTree[i] != wantedDirTree[i] {
// 			t.Errorf("Trees don't match")
// 		}
// 	}

// 	c = configs{
// 		perm: true,
// 	}
// 	gotPermTree := getDirTree(dirPath, c)

// 	wantedPermTree := []string{
// 		"[-rwxrwxr-x]resources",
// 		"│\t│──[-rw-rw-r--]sample1.txt",
// 		"│\t│──[-rwxrwxr-x]test-folder1",
// 		"│\t│\t│──[-rwxrwxr-x]test-folder2",
// 		"│\t│\t│\t│──[-rw-rw-r--]sample.txt",
// 	}
// 	for i:=0; i<len(wantedPermTree); i++ {
// 		if gotPermTree[i] != wantedPermTree[i] {
// 			t.Errorf("Trees don't match")
// 		}
// 	}

// 	c = configs{
// 		level: 2
// 	}
// 	gotLevel2Tree := getDirTree(dirPath, c)

// 	wantedLevel2Tree := []string{
// 		"resources",
// 		"│\t│──sample1.txt",
// 		"│\t│──test-folder1",
// 	}
// 	for i:=0; i<len(wantedLevel2Tree); i++ {
// 		if gotLevel2Tree[i] != wantedLevel2Tree[i] {
// 			t.Errorf("Trees don't match")
// 		}
// 	}

// 	// deleting the created dir
// 	os.RemoveAll("./resources/")

// }
