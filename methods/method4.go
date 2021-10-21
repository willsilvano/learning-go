package main

import (
	"fmt"
	"methods/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Permiter:", t.Perimeter())
}
