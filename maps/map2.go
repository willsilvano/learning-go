package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["Doe"] = 30
	fmt.Println(studentsAge)
}
