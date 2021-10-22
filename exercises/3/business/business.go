package business

import "fmt"

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Credits float32
	Account
}

func (a *Account) ChangeName(firstName string, lastName string) {
	a.FirstName = firstName
	a.LastName = lastName
}

func (a Account) String() string {
	return fmt.Sprintf("FirstName: %s, LastName: %s", a.FirstName, a.LastName)
}

func (e *Employee) AddCredits(credit float32) {
	e.Credits += credit
}

func (e *Employee) RemoveCredits(credit float32) {
	e.Credits -= credit
}

func (e Employee) CheckCredits() float32 {
	return e.Credits
}

func (e Employee) String() string {
	return fmt.Sprintf("%s, Credits: %f", e.Account, e.Credits)
}
