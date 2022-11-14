package main

import (
	"fmt"
)

func oddOrEven(x int) string {
	if x%2 == 0 {
		return "even"
	} else if x%2 == 1 {
		return "odd"
	} else {
		return "invalid"
	}

}

func main() {
	fmt.Println(oddOrEven(-4), oddOrEven(7))
}
