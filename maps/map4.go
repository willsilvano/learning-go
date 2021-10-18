package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["Doe"] = 31
	fmt.Println("Maria's age is", studentsAge["Maria"])
}
