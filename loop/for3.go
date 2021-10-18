package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Println("Writting inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("Finish!")
			break
		}
		fmt.Println(num)
	}
}
