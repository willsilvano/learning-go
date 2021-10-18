package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["Doe"] = 31

	age, exist := studentsAge["Maria"]
	if exist {
		fmt.Println("Maria's age is", age)
	} else {
		fmt.Println("Maria's age couldn't be found")
	}
}
