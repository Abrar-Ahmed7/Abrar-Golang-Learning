package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p Point) Offset(x, y float64) Point {
	p.X = p.X + x
	p.Y = p.Y + y
	return Point{p.X, p.Y}
}

func (p *Point) Move(x, y float64) Point {
	p.X = p.X + x
	p.Y = p.Y + y
	return Point{p.X, p.Y}
}

func main() {
	V := Point{2, 5}
	fmt.Printf("Before using Offset(): %v \n", V)
	fmt.Printf("Calling Offset(): %v \n", V.Offset(2, 3))
	fmt.Printf("After using Offset(): %v \n", V)
	P := &V
	fmt.Printf("Before using Move(): %v \n", *P)
	fmt.Printf("Calling Move(): %v \n", P.Move(2, 3))
	fmt.Printf("After using Move(): %v \n", *P)
}
