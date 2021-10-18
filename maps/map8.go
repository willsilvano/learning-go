package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["Doe"] = 31

	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
}
