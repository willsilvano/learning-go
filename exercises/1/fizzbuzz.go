package main

import "fmt"

func fizzbuss(number int) {
	switch {
	case number%3 == 0 && number%5 == 0:
		fmt.Println("FizzBuzz")
	case number%3 == 0:
		fmt.Println("Fizz")
	case number%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(number)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fizzbuss(i)
	}
}
