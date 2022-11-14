package main

import "fmt"

func main() {
	var myArray [3]string
	myArray[0] = "Dummy"
	fmt.Println(myArray)
	var numbers = [3]int{10, 20, 30}
	fmt.Println(numbers)
	var sparsedArray = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(sparsedArray)
}
