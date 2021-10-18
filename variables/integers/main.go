package main

import "fmt"

func main() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)

	// var integer32Two int32 = 2147483648
	// fmt.Println(integer32Two)

	// var challengeTwo uint = -10
	// fmt.Println(challengeTwo)

}
