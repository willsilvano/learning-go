package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["Doe"] = 31

	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
