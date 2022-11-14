package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (n *List) AddNode(data T) {
	newNode := List{data, nil}
	iter := n
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = &newNode
}

func (n *List) PrintNode() {
	iter := n
	for iter != nil {
		fmt.Println(iter.val)
		iter = iter.next
	}
}

func main() {
	newNode := make(List[int], 0)
	newNode = List{nil, 10}
	newNode.AddNode(20)
	newNode.AddNode(30)
	newNode.AddNode(40)
	newNode.PrintNode()
}
