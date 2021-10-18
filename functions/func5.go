package main

import "fmt"

func main() {
	myName := "Willian"
	updateName(myName)
	fmt.Println(myName)
}

func updateName(name string) {
	name = "Teste"
}
