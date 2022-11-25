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
				"      └──sample.txt\n" +
				"\n" +
				"3 directories, 2 files",
		},
		{
			dirPath: "./resources",
			input:   configs{dirOnly: true},
			r:       report{},
			want: "resources\n" +
				"└──test-folder1\n" +
				"   └──test-folder2\n" +
				"\n" +
				"2 directories",
		},
		{
			dirPath: "./resources",
			input:   configs{level: 1},
			r:       report{},
			want: "resources\n" +
				"│──sample1.txt\n" +
				"└──test-folder1\n" +
				"\n" +
				"2 directories, 1 files",
		},
	}

	for _, tc := range tests {
		tree, _ := tree(tc.dirPath, "", "", "", &tc.r, tc.input, 0)
		got := appendReport(tree, tc.r, tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected:\n %v, got:\n %v", tc.want, got)
		}
	}

	// deleting the created dir
	os.RemoveAll("./resources/")
}

func TestJsonTree(t *testing.T) {
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
			input:   configs{json: true},
			r:       report{},
			want: "[\n" +
				"  {\"type\":\"directory\",\"name\":\"resources\",\"contents\":[\n" +
				"    {\"type\":\"file\",\"name\":\"sample1.txt\"},\n" +
				"    {\"type\":\"directory\",\"name\":\"test-folder1\",\"contents\":[\n" +
				"      {\"type\":\"directory\",\"name\":\"test-folder2\",\"contents\":[\n" +
				"        {\"type\":\"file\",\"name\":\"sample.txt\"}\n" +
				"      ]}\n" +
				"    ]}\n" +
				"  ]}\n" +
				",\n" +
				"  {\"type\":\"report\",\"directories\":3, \"files\":2}\n" +
				"]",
		},
		// {
		// 	dirPath: "./resources",
		// 	input:   configs{dirOnly: true},
		// 	r:       report{},
		// 	want: "[\n" +
		// 	"  {\"type\":\"directory\",\"name\":\"resources\",\"contents\":[\n" +
		// 	"    {\"type\":\"file\",\"name\":\"sample1.txt\"},\n" +
		// 	"    {\"type\":\"directory\",\"name\":\"test-folder1\",\"contents\":[\n" +
		// 	"      {\"type\":\"directory\",\"name\":\"test-folder2\",\"contents\":[\n" +
		// 	"        {\"type\":\"file\",\"name\":\"sample.txt\"}\n" +
		// 	"      ]}\n" +
		// 	"    ]}\n" +
		// 	"  ]}\n" +
		// 	",\n" +
		// 	"  {\"type\":\"report\",\"directories\":2, \"files\":2}\n" +
		// 	"]",
		// },
		{
			dirPath: "./resources",
			input:   configs{json: true, level: 1},
			r:       report{},
			want: "[\n" +
				"  {\"type\":\"directory\",\"name\":\"resources\",\"contents\":[\n" +
				"    {\"type\":\"file\",\"name\":\"sample1.txt\"},\n" +
				"    {\"type\":\"directory\",\"name\":\"test-folder1\"  ]}\n" +
				",\n" +
				"  {\"type\":\"report\",\"directories\":2, \"files\":1}\n" +
				"]",
		},
	}

	for _, tc := range tests {
		tree, _ := jsonTree(tc.dirPath, "  ", "", &tc.r, tc.input, 1)
		got := addSqBrkt(appendReport(tree, tc.r, tc.input))
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected:\n %v, got:\n %v", tc.want, got)
		}
	}

	// deleting the created dir
	os.RemoveAll("./resources/")

}

func TestXmlTree(t *testing.T) {
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
			input:   configs{xml: true},
			r:       report{},
			want: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<tree>\n" +
				"  <directory name=\"resources\">\n" +
				"	<file name=\"sample1.txt\"></file>\n" +
				"	<directory name=\"test-folder1\">\n" +
				"	  <directory name=\"test-folder2\">\n" +
				"		<file name=\"sample.txt\"></file>\n" +
				"	  </directory>\n" +
				"	</directory>\n" +
				"  </directory>\n" +
				"  <report>\n    <directories>" +
				"3</directories>\n	<files>2</files>" +
				"\n  </report>" +
				"\n</tree>",
		},
		{
			dirPath: "./resources",
			input:   configs{xml: true, dirOnly: true},
			r:       report{},
			want: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<tree>\n" +
				"  <directory name=\"resources\">\n" +
				"	<directory name=\"test-folder1\">\n" +
				"	  <directory name=\"test-folder2\">\n" +
				"	  </directory>\n" +
				"	</directory>\n" +
				"  <report>\n    <directories>" +
				"2</directories>" +
				"\n  </report>" +
				"\n</tree>",
		},
		{
			dirPath: "./resources",
			input:   configs{xml: true, level: 1},
			r:       report{},
			want: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<tree>\n" +
				"  <directory name=\"resources\">\n" +
				"	<file name=\"sample1.txt\"></file>\n" +
				"	<directory name=\"test-folder1\">  </directory>" +
				"  <report>\n    <directories>" +
				"2</directories>\n	<files>1</files>" +
				"\n  </report>" +
				"\n</tree>",
		},
	}

	for _, tc := range tests {
		tree, _ := xmlTree(tc.dirPath, "  ", "", &tc.r, tc.input, 1)
		got := addTreeTag(appendReport(tree, tc.r, tc.input))
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected:\n %v, got:\n %v", tc.want, got)
		}
	}

	// deleting the created dir
	os.RemoveAll("./resources/")

}
