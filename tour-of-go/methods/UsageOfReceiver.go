package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(bytes []byte) (int, error) {
	l := len(bytes)
	*b += ByteCounter(l)
	return l, nil
}

func main() {

	var counter ByteCounter
	f1, _ := os.Open("temp.txt")
	// f2, _ := os.Create("dest.txt")
	f2 := &counter
	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
	fmt.Println(counter)

}
