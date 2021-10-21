package main

import (
	"fmt"
)

func fibonacci(number int) []int {
	if number < 2 {
		return make([]int, 0)
	}

	fibonacciSequence := make([]int, number)
	fibonacciSequence[0], fibonacciSequence[1] = 1, 1

	for i := 2; i < number; i++ {
		fibonacciSequence[i] = fibonacciSequence[i-1] + fibonacciSequence[i-2]
	}

	return fibonacciSequence
}

func main() {

	var number int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number)

	fmt.Println("Number:", number)
	fmt.Println("Fibonacci sequence:", fibonacci(number))
}
