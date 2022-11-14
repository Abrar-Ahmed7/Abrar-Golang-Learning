package main

import "fmt"

type I interface {
	printString()
}

type T struct {
	S string
}

func (t T) printString() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.printString()
}
