package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Permiter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Permiter() float64 {
	return s.size * 4
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Permiter() float64 {
	return 2 * math.Pi * c.radius
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Permiter: ", s.Permiter())
	fmt.Println()
}

func main() {
	var s Shape = Square{3}
	printInformation(s)

	c := Circle{6}
	printInformation(c)
}
