package main

import (
	"exercises/3/business"
	"fmt"
)

func main() {
	account := business.Account{
		FirstName: "Will",
		LastName:  "SM",
	}

	employee := business.Employee{
		Credits: 0,
		Account: account,
	}

	fmt.Println(employee)
	employee.ChangeName("Willian", "Silvano Maria")
	fmt.Println(employee)
	employee.AddCredits(100)
	fmt.Println(employee)
	employee.AddCredits(50)
	fmt.Println(employee)
	fmt.Println(employee.CheckCredits())

}
