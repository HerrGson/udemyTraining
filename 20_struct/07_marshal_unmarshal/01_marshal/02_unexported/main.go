package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first    string
	last     string
	age      int
	Exported int
}

func main() {
	p1 := person{"James", "Bond", 20, 01}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
