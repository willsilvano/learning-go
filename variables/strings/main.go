package main

import "fmt"

func main() {
	var firstName string = "Willian"
	lastName := "Maria"

	fmt.Println(firstName, lastName)

	fmt.Println("\nTeste1\nTeste2")
	fmt.Println("\tTeste1\tTeste2")
	fmt.Println("\tTeste\"1\tTeste\"2")
	fmt.Println("\tTeste\\Teste2")

	fullName := "\n\nWillian Maria \t(alias \"Will\")\n"
	fmt.Println(fullName)
}
