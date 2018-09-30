package main

import "fmt"

func main() {
	number := 3
	fmt.Println(calc(number))

	funcExpression := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(funcExpression(number))
}

func calc(n int) (quotient int, even bool) {
	quotient = n / 2
	even = n%2 == 0
	return
}
