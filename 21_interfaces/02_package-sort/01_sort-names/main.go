package main

import (
	"fmt"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func Swap(i, j int, p people) {
	p[i], p[j] = p[j], p[j]
}

func main() {
	// studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	studyGroup2 := people{"Zeno", "John", "Al", "Jenny"}

	// fmt.Println(studyGroup)
	// sort.Sort(studyGroup)
	// fmt.Println(studyGroup)

	fmt.Println(studyGroup2)
	Swap(0, 1, studyGroup2)
	fmt.Println(studyGroup2)

}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface
