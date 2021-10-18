package main

import "fmt"

func highLow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highLow() low greater than high")
	}
	defer fmt.Println("Deferred: highLow(", high, ",", low, ")")
	fmt.Println("Call: highLow(", high, ",", low, ")")

	highLow(high, low+1)
}

func main() {
	highLow(2, 0)
	fmt.Println("Program finished successfully!")
}
