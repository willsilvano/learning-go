package main

import "fmt"

func main() {
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}

	fmt.Println("Cities:", cities)
	fmt.Println(cities[0])
	fmt.Println(cities[len(cities)-1])
}
