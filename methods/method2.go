package main

import (
	"fmt"
	"strings"
)

type upperString string

func (u upperString) Upper() string {
	return strings.ToUpper(string(u))
}

func main() {
	s := upperString("Learning Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())
}
