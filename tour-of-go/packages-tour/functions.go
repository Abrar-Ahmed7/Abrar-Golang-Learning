package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func newAdd(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(add(12, 122))
	fmt.Println(newAdd(23, 45))
}
