package main

import "fmt"

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func main() {
	t := triangle{size: 3}
	s := square{size: 6}
	fmt.Println("Size (triangle):", t.size)
	fmt.Println("Perimeter (triangle):", t.perimeter())
	t.doubleSize()
	fmt.Println("Size (triangle):", t.size)
	fmt.Println("Perimeter (triangle):", t.perimeter())

	fmt.Println("Perimeter (square):", s.perimeter())
}
