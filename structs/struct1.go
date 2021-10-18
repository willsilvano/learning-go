package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	// john := Employee{1, "John", "Doe", "Doe's Street"}
	john := Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Address:   "Doe's Street",
	}

	fmt.Println(john)
	fmt.Printf("FirstName: %s\nLastName: %s\nAddress: %s", john.FirstName, john.LastName, john.Address)
}
