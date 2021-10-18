package main

import "fmt"

func main() {
	myName := "Willian"
	updateName2(&myName)
	fmt.Println(myName)
}

func updateName2(name *string) {
	*name = "Teste"
}
