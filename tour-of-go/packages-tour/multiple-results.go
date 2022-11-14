package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

/*
Here x and y will be returned
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("first", "second")
	fmt.Println(a, b)
	fmt.Println(split(22))

	// type- inference
	v := -42.12
	fmt.Printf("v is of type %T\n", v)
}
