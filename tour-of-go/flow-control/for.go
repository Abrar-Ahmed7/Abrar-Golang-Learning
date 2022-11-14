package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//with condition only also this is used as while in go
	num := 1
	for num < 1000 {
		num += num
	}
	fmt.Println(num)
}
