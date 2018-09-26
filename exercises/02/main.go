package main

import (
	"fmt"
)

func main() {
	var small int
	var big int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Please enter a big number: ")
	fmt.Scan(&big)

	remainder := big % small
	fmt.Printf("Smal number is %d and the big one is %d \n and the remainder is %d", small, big, remainder)
}
