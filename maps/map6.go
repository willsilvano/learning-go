package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["Doe"] = 31

	delete(studentsAge, "Doe")
	fmt.Println(studentsAge)
}
