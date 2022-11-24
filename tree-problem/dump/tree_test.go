package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	// creating dir
	if err := os.MkdirAll("resources/test-folder1/test-folder2", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	_, err1 := os.Create("resources/test-folder1/test-folder2/sample.txt")
	if err1 != nil {
		fmt.Printf("Unable to open file: %s", err1)
	}
	_, err2 := os.Create("resources/sample1.txt")
	if err2 != nil {
		fmt.Printf("Unable to open file: %s", err2)
	}

	tests := []struct {
		dirPath string
		input   configs
		r       report
		want    string
	}{
		{
			dirPath: "./resources",
			input:   configs{},
			r:       report{},
			want: "resources\n" +
				"│──sample1.txt\n" +
				"└──test-folder1\n" +
				"   └──test-folder2\n" +
				"      └──sample.txt\n",
		},
		{
			dirPath: "./resources",
			input:   configs{dirOnly: true},
			r:       report{},
			want: "resources\n" +
				"└──test-folder1\n" +
				"   └──test-folder2\n",
		},
		{
			dirPath: "./resources",
			input:   configs{level: 1},
			r:       report{},
			want: "resources\n" +
				"│──sample1.txt\n" +
				"└──test-folder1\n",
		},
	}

	for _, tc := range tests {
		got, _ := tree(tc.dirPath, "", "", "", &tc.r, tc.input, 0)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected:\n %v, got:\n %v", tc.want, got)
		}
	}

	// deleting the created dir
	os.RemoveAll("./resources/")
}
