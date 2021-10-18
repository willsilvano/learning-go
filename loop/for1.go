package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println("sum of 1..100 is", sum)
}
