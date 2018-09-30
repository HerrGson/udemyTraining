package main

import "fmt"

var x int

func hello() {
	fmt.Print("hello ")
	x++
}

func world() {
	fmt.Println("world")
	x += 2
}

func main() {
	defer world()
	fmt.Println(x)
	hello()
	fmt.Println(x)
}
