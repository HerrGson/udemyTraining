package main

import "fmt"

func greatest(a ...int) int {
	var max int
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println(greatest(1, 3, 7, 2, 5, 9))
}
