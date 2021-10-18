package main

import "fmt"

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter2 := months[3:6]
	quarter2Extended := quarter2[:4]
	quarter2Extended2 := quarter2[0:]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
	fmt.Println(quarter2Extended2, len(quarter2Extended2), cap(quarter2Extended2))
}
