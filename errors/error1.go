package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("Employee not found!")

func main() {
	employee, err := getInformation(1001)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND: %v", err)
	} else {
		fmt.Print(employee)
	}
}

func getInformation(id int) (*Employee, error) {

	if id != 1001 {
		// return nil, fmt.Errorf("Got an error whe getting the employee information: %v", err)
		return nil, ErrNotFound
	}

	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

// func apiCallEmployee(id int) (*Employee, error) {
// 	employee := Employee{LastName: "Doe", FirstName: "John"}
// 	return &employee, nil
// }
